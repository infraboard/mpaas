package build

import (
	"github.com/infraboard/mcenter/common/validate"
	request "github.com/infraboard/mcube/http/request"
)

const (
	AppName = "build"
)

type Service interface {
	RPCServer
}

func (req *CreateBuildConfigRequest) Validate() error {
	return validate.Validate(req)
}

func NewQueryBuildConfigRequest() *QueryBuildConfigRequest {
	return &QueryBuildConfigRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
