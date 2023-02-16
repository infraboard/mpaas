package approval

import (
	"encoding/json"
	"time"

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
	req.Proposers = append(req.Proposers, userIds...)
}

func (req *CreateApprovalRequest) AddAuditor(userIds ...string) {
	req.Auditors = append(req.Auditors, userIds...)
}

func (req *CreateApprovalRequest) UserIds() (uids []string) {
	uids = append(uids, req.Auditors...)
	uids = append(uids, req.Proposers...)
	return
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
		Spec: &CreateApprovalRequest{},
	}
}

func (i *Approval) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateApprovalRequest
		*Status
		DeployPipeline *pipeline.Pipeline
	}{i.Meta, i.Spec, i.Status, i.DeployPipeline})
}

func (s *Status) IsAllowPublish() bool {
	if s.Stage > STAGE_PASSED && s.Stage < STAGE_CLOSED {
		return true
	}

	return false
}

func (s *Status) Update(stage STAGE) {
	s.Stage = stage
	switch stage {
	case STAGE_DENY, STAGE_PASSED:
		s.AuditAt = time.Now().Unix()
	case STAGE_PUBLISHING:
		s.PublishAt = time.Now().Unix()
	case STAGE_CANCELED:
		s.CancelAt = time.Now().Unix()
	case STAGE_SUCCEEDED, STAGE_FAILED:
		s.EndAt = time.Now().Unix()
	case STAGE_CLOSED:
		s.CloseAt = time.Now().Unix()
	}
}
