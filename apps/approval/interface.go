package approval

import (
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "approvals"
)

type Service interface {
	RPCServer
}

func NewQueryApprovalRequest() *QueryApprovalRequest {
	return &QueryApprovalRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *CreateApprovalRequest) Validate() error {
	return validate.Validate(req)
}

func NewCreateApprovalRequest() *CreateApprovalRequest {
	return &CreateApprovalRequest{}
}

func (req *DescribeApprovalRequest) Validate() error {
	return validate.Validate(req)
}

func (req *EditApprovalRequest) Validate() error {
	return validate.Validate(req)
}

func NewUpdateApprovalStatusRequest(approvalId string) *UpdateApprovalStatusRequest {
	status := NewStatus()
	return &UpdateApprovalStatusRequest{
		Id:     approvalId,
		Status: status,
	}
}

func (req *UpdateApprovalStatusRequest) Validate() error {
	return validate.Validate(req)
}

func NewDescribeApprovalRequest(id string) *DescribeApprovalRequest {
	return &DescribeApprovalRequest{
		Id: id,
	}
}
