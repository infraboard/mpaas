package job

import (
	"github.com/infraboard/mcenter/common/validate"
	request "github.com/infraboard/mcube/http/request"
)

const (
	AppName = "jobs"
)

type Service interface {
	RPCServer
}

func (req *CreateJobRequest) Validate() error {
	return validate.Validate(req)
}

func NewCreateJobRequest() *CreateJobRequest {
	return &CreateJobRequest{}
}

func NewQueryJobRequest() *QueryJobRequest {
	return &QueryJobRequest{
		Page:       request.NewDefaultPageRequest(),
		WithGlobal: true,
		Ids:        []string{},
		Names:      []string{},
	}
}
