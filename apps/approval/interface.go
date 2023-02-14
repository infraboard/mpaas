package approval

import "github.com/infraboard/mcube/http/request"

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
