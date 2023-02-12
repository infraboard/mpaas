package pipeline

import (
	"fmt"
	"time"

	"github.com/infraboard/mcenter/common/validate"
	pb_request "github.com/infraboard/mcube/pb/request"
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

func NewPutPipelineRequest(id string) *UpdatePipelineRequest {
	return &UpdatePipelineRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().Unix(),
		Spec:       NewCreatePipelineRequest(),
	}
}

func NewPatchPipelineRequest(id string) *UpdatePipelineRequest {
	return &UpdatePipelineRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().Unix(),
		Spec:       NewCreatePipelineRequest(),
	}
}

func NewDeletePipelineRequest(id string) *DeletePipelineRequest {
	return &DeletePipelineRequest{
		Id: id,
	}
}
