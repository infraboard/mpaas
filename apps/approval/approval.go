package approval

import "github.com/infraboard/mpaas/common/meta"

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
