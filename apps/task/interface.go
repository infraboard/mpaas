package task

import (
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

func NewRunTaskRequest(spec string, params *job.VersionedRunParam) *RunTaskRequest {
	return &RunTaskRequest{
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
