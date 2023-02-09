package pipeline

import (
	"fmt"

	"github.com/infraboard/mcenter/common/validate"
)

const (
	AppName = "pipelines"
)

type Service interface {
	RPCServer
}

func NewDescribePipelineRequest(id string) *DescribePipelineRequest {
	return &DescribePipelineRequest{
		Id: id,
	}
}

func (req *DescribePipelineRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("Pipeline Id required")
	}
	return nil
}

func (req *CreatePipelineRequest) Validate() error {
	return validate.Validate(req)
}
