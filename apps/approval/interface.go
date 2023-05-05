package approval

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mpaas/apps/job"
)

const (
	AppName = "approvals"
)

type Service interface {
	// 删除发布申请
	DeleteApproval(context.Context, *DeleteApprovalRequest) (*Approval, error)
	RPCServer
}

func NewDeleteApprovalRequest(id string) *DeleteApprovalRequest {
	return &DeleteApprovalRequest{
		Id: id,
	}
}

type DeleteApprovalRequest struct {
	Id string
}

func NewQueryApprovalRequest() *QueryApprovalRequest {
	return &QueryApprovalRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *CreateApprovalRequest) Validate() error {
	if req.PipelineSpec == nil &&
		req.PipelineId == "" {
		return fmt.Errorf("流水线配置缺失")
	}

	if len(req.Operators) == 0 {
		return fmt.Errorf("执行人缺失")
	}

	if len(req.Auditors) == 0 {
		return fmt.Errorf("审核人缺失")
	}

	return validate.Validate(req)
}

func NewCreateApprovalRequest() *CreateApprovalRequest {
	return &CreateApprovalRequest{
		Operators: []string{},
		Auditors:  []string{},
		RunParams: []*job.RunParam{},
		Labels:    map[string]string{},
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

func NewQueryApprovalRequestFromHTTP(r *http.Request) *QueryApprovalRequest {
	return &QueryApprovalRequest{
		Page: request.NewPageRequestFromHTTP(r),
	}
}
