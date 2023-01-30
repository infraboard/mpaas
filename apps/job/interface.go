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
	return &CreateJobRequest{
		RunnerParams: make(map[string]string),
		RunParams:    []*VersionedRunParamDesc{},
		Labels:       make(map[string]string),
	}
}

func NewQueryJobRequest() *QueryJobRequest {
	return &QueryJobRequest{
		Page:  request.NewDefaultPageRequest(),
		Ids:   []string{},
		Names: []string{},
	}
}

func NewDescribeJobRequest(id string) *DescribeJobRequest {
	return &DescribeJobRequest{
		DescribeValue: id,
	}
}

func (req *DescribeJobRequest) Validate() error {
	return validate.Validate(req)
}
