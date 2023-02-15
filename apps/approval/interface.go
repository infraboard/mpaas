package approval

import (
	"fmt"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mpaas/apps/pipeline"
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
	if req.DeployPipelineSpec == nil &&
		req.DeployPipelineId == "" {
		return fmt.Errorf("流水线配置缺失")
	}

	if len(req.Proposers) == 0 {
		return fmt.Errorf("申请人缺失")
	}

	if len(req.Auditors) == 0 {
		return fmt.Errorf("审核人缺失")
	}

	return validate.Validate(req)
}

func NewCreateApprovalRequest() *CreateApprovalRequest {
	return &CreateApprovalRequest{
		DeployPipelineSpec: pipeline.NewCreatePipelineRequest(),
	}
}

func (req *DescribeApprovalRequest) Validate() error {
	return validate.Validate(req)
}

func NewEditApprovalRequest(approvalId string) *EditApprovalRequest {
	return &EditApprovalRequest{
		Id: approvalId,
	}
}

func (req *EditApprovalRequest) Validate() error {
	return validate.Validate(req)
}

func NewUpdateApprovalStatusRequest(approvalId string) *UpdateApprovalStatusRequest {
	return &UpdateApprovalStatusRequest{
		Id:     approvalId,
		Status: NewStatus(),
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
