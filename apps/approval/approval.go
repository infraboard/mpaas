package approval

import (
	"time"

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

func (s *Status) Update(stage STAGE) {
	s.Stage = stage
	switch stage {
	case STAGE_DENY, STAGE_PASSED:
		s.AuditAt = time.Now().Unix()
	case STAGE_PUBLISHING:
		s.PublishAt = time.Now().Unix()
	case STAGE_SUCCEEDED, STAGE_FAILED:
		s.EndAt = time.Now().Unix()
	case STAGE_CLOSED:
		s.CloseAt = time.Now().Unix()
	}
}
