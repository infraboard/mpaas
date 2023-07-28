package pipeline

import (
	"fmt"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	job "github.com/infraboard/mpaas/apps/job"
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

func (req *CreatePipelineRequest) GetLabelValue(key string) string {
	return req.Labels[key]
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

func (req *CreatePipelineRequest) AddLabel(key, value string) {
	req.Labels[key] = value
}

func NewDeletePipelineRequest(id string) *DeletePipelineRequest {
	return &DeletePipelineRequest{
		Id: id,
	}
}

func NewQueryPipelineRequestFromHTTP(r *restful.Request) *QueryPipelineRequest {
	req := NewQueryPipelineRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	req.Filters = policy.GetScopeFilterFromRequest(r)
	isTemp := r.QueryParameter("is_template")
	if isTemp != "" {
		req.SetIsTemplate(isTemp == "true")
	}
	return req
}

func NewQueryPipelineRequest() *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (r *QueryPipelineRequest) SetIsTemplate(v bool) {
	r.IsTemplate = &v
}

func NewRunPipelineRequest(pipelineId string) *RunPipelineRequest {
	return &RunPipelineRequest{
		PipelineId: pipelineId,
		RunParams:  []*job.RunParam{},
		Labels:     make(map[string]string),
	}
}
