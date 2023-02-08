package task

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
)

const (
	AppName = "tasks"
)

type Service interface {
	JobService
	PipelineService
}

type JobService interface {
	JobRPCServer
}

func NewQueryTaskRequest() *QueryJobTaskRequest {
	return &QueryJobTaskRequest{
		Page: request.NewDefaultPageRequest(),
		Ids:  []string{},
	}
}

func NewRunTaskRequest(name, spec string, params *job.VersionedRunParam) *RunTaskRequest {
	return &RunTaskRequest{
		Name:    name,
		JobSpec: spec,
		Params:  params,
	}
}

type PipelineService interface {
	PipelineRPCServer
}

func NewRunPipelineRequest() *RunPipelineRequest {
	return &RunPipelineRequest{}
}

func NewDescribeJobTaskRequest(id string) *DescribeJobTaskRequest {
	return &DescribeJobTaskRequest{
		Id: id,
	}
}

func (r *DescribeJobTaskRequest) Validate() error {
	if r.Id == "" {
		return fmt.Errorf("job id required")
	}
	return nil
}

func NewDeleteJobTaskRequest(id string) *DeleteJobTaskRequest {
	return &DeleteJobTaskRequest{
		Id: id,
	}
}
