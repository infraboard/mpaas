package approval

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/infraboard/mpaas/apps/approval/impl/notify"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/common/meta"
)

func New(req *CreateApprovalRequest) (*Approval, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &Approval{
		Meta:   meta.NewMeta(),
		Spec:   req,
		Status: NewStatus(),
	}, nil
}

func (req *CreateApprovalRequest) AddProposer(userIds ...string) {
	req.Operators = append(req.Operators, userIds...)
}

func (req *CreateApprovalRequest) AddAuditor(userIds ...string) {
	req.Auditors = append(req.Auditors, userIds...)
}

func (req *CreateApprovalRequest) UserIds() (uids []string) {
	uids = append(uids, req.Auditors...)
	uids = append(uids, req.Operators...)
	return
}

func (req *CreateApprovalRequest) AutoRunDesc() string {
	if req.AutoRun {
		return "自动执行"
	}
	return "手动执行"
}

func (req *CreateApprovalRequest) RunParamsDesc() string {
	buf := bytes.NewBufferString("\\n")
	for i := range req.RunParams {
		item := req.RunParams[i]
		buf.WriteString(item.MarkdownShortShow())
		buf.WriteString("\\n")
	}
	return buf.String()
}

func (req *CreateApprovalRequest) IsAuditor(uid string) bool {
	for _, v := range req.Auditors {
		if v == uid {
			return true
		}
	}
	return false
}

func NewStatus() *Status {
	return &Status{}
}

func NewApprovalSet() *ApprovalSet {
	return &ApprovalSet{
		Items: []*Approval{},
	}
}

func (s *ApprovalSet) Add(item *Approval) {
	s.Items = append(s.Items, item)
}

func NewDefaultApproval() *Approval {
	return &Approval{
		Meta: meta.NewMeta(),
		Spec: NewCreateApprovalRequest(),
	}
}

func (i *Approval) UUID() string {
	return fmt.Sprintf("approval-%s", i.Meta.Id)
}

func (i *Approval) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateApprovalRequest
		*Status
		DeployPipeline *pipeline.Pipeline
	}{i.Meta, i.Spec, i.Status, i.Pipeline})
}

func (i *Approval) FeishuAuditNotifyMessage() *notify.FeishuAuditNotifyMessage {
	msg := notify.NewFeishuAuditNotifyMessage()
	msg.Domain = i.Spec.Domain
	msg.Namespace = i.Spec.Namespace
	msg.ApprovalId = i.Meta.Id
	msg.Title = i.Spec.Title
	msg.CreateBy = i.Spec.CreateBy
	// msg.Operator = i.Spec.Operators[]
	msg.PipelineDesc = i.Pipeline.Spec.Description
	msg.ExecType = i.Spec.AutoRunDesc()
	msg.ExecVars = i.Spec.RunParamsDesc()
	msg.DenyButtonName = "拒绝"
	msg.PassButtonName = "同意"
	msg.Note = "该消息由mpaas平台提供"

	switch i.Status.Stage {
	// 待审核, 通知审核人
	case STAGE_PENDDING:
		msg.ShowDenyButton = true
		msg.ShowPassButton = true
		// msg.Auditor = ""
	// 审核通过, 通知申请人, 通知其他审核人
	case STAGE_PASSED:
		msg.ShowPassButton = true
		msg.PassButtonName = "xxx已同意"
		// msg.Auditor = ""
	// 审核通过, 通知申请人, 通知其他审核人
	case STAGE_DENY:
		msg.ShowDenyButton = true
		msg.PassButtonName = "xxx已拒绝"
	// 审核过期, 通知所有人
	case STAGE_EXPIRED:
		msg.ShowPassButton = true
		msg.PassButtonName = "已过期"
	// 审核关闭, 通知所有人
	case STAGE_CLOSED:
		msg.ShowPassButton = true
		msg.PassButtonName = "已关闭"
	}
	return msg
}

func (s *Status) IsAllowPublish() bool {
	if s.Stage >= STAGE_PASSED && s.Stage < STAGE_CLOSED {
		return true
	}

	return false
}

func (s *Status) AddNotifyRecords(records ...*NotifyRecord) {
	s.NotifyRecords = append(s.NotifyRecords, records...)
}

func (s *Status) Update(stage STAGE) {
	s.Stage = stage
	switch stage {
	case STAGE_DENY, STAGE_PASSED:
		s.AuditAt = time.Now().Unix()
	case STAGE_CLOSED:
		s.CloseAt = time.Now().Unix()
	}
}

func NewNotifyRecord(Stage STAGE) *NotifyRecord {
	return &NotifyRecord{
		Stage:    Stage,
		NotifyAt: time.Now().Unix(),
	}
}

func (r *NotifyRecord) Failed(err error) {
	r.Message = err.Error()
}

func (r *NotifyRecord) Success(detail string) {
	r.IsSuccess = true
	r.Detail = detail
}
