package approval

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mpaas/apps/approval/impl/notify"
	job "github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
)

func New(req *CreateApprovalRequest) (*Approval, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Approval{
		Meta:   resource.NewMeta().IdWithPrefix("apv"),
		Spec:   req,
		Status: NewStatus(),
	}, nil
}

func (req *CreateApprovalRequest) UpdateFromToken(tk *token.Token) {
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
	req.CreateBy = tk.UserId
}

func (req *CreateApprovalRequest) AddOperator(userIds ...string) {
	req.Operators = append(req.Operators, userIds...)
}

func (req *CreateApprovalRequest) OperatorToString() string {
	return strings.Join(req.Operators, ",")
}

func (req *CreateApprovalRequest) AddAuditor(userIds ...string) {
	req.Auditors = append(req.Auditors, userIds...)
}

func (req *CreateApprovalRequest) AuditorToString() string {
	return strings.Join(req.Auditors, ",")
}

func (req *CreateApprovalRequest) AddRunParam(params ...*job.RunParam) {
	req.RunParams = append(req.RunParams, params...)
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
		Meta: resource.NewMeta(),
		Spec: NewCreateApprovalRequest(),
	}
}

// 其他审核人
func (i *Approval) OtherAuditors() (users []string) {
	for _, auditor := range i.Spec.Auditors {
		if auditor != i.Status.AuditBy {
			users = append(users, auditor)
		}
	}
	return
}

// 操作人和其他审核人
func (i *Approval) OperatorAndOtherAuditors() (users []string) {
	users = append(users, i.Spec.Operators...)
	users = append(users, i.OtherAuditors()...)
	return
}

// 所有人
func (i *Approval) OperatorAndAuditors() (users []string) {
	users = append(users, i.Spec.Operators...)
	users = append(users, i.Spec.Auditors...)
	return
}

func (i *Approval) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateApprovalRequest
		*Status
		Pipeline *pipeline.Pipeline `json:"pipeline"`
	}{i.Meta, i.Spec, i.Status, i.Pipeline})
}

func (i *Approval) FeishuAuditNotifyMessage() (msg *notify.FeishuAuditNotifyMessage, users []string) {
	msg = notify.NewFeishuAuditNotifyMessage()
	msg.Domain = i.Spec.Domain
	msg.Namespace = i.Spec.Namespace
	msg.ApprovalId = i.Meta.Id
	msg.Title = i.Spec.Title
	msg.CreateBy = i.Spec.CreateBy
	msg.Operator = i.Spec.OperatorToString()
	msg.Auditor = i.Spec.AuditorToString()
	msg.PipelineDesc = i.Pipeline.Spec.Description
	msg.ExecType = i.Spec.AutoRunDesc()
	msg.ExecVars = i.Spec.RunParamsDesc()
	msg.DenyButtonName = "拒绝"
	msg.PassButtonName = "同意"
	msg.Note = "该消息由mpaas平台提供"

	switch i.Status.Stage {
	case STAGE_PENDDING:
		// 待审核, 通知审核人
		users = i.Spec.Auditors
	case STAGE_PASSED:
		// 审核通过, 通知申请人, 通知其他审核人
		msg.ShowDenyButton = false
		msg.PassButtonName = i.Status.AuditBy + "已同意"
		users = i.OperatorAndOtherAuditors()
	case STAGE_DENY:
		// 审核通过, 通知申请人, 通知其他审核人
		msg.ShowDenyButton = false
		msg.PassButtonName = i.Status.AuditBy + "已拒绝"
		users = i.OperatorAndOtherAuditors()
	case STAGE_EXPIRED:
		// 审核过期, 通知所有人
		msg.ShowDenyButton = false
		msg.PassButtonName = "已过期"
		users = i.OperatorAndAuditors()
	case STAGE_CLOSED:
		// 审核关闭, 通知所有人
		msg.ShowDenyButton = false
		msg.PassButtonName = "已关闭"
		users = i.OperatorAndAuditors()
	}

	return
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

// 根据申请单状态判断是否可以删除, 草稿状态和关闭状态的才允许删除
func (s *Status) AllowDelete() bool {
	return s.Stage < STAGE_PENDDING || s.Stage >= STAGE_CLOSED
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
