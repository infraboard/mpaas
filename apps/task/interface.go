package task

import (
	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
)

const (
	AppName = "tasks"
)

type Service interface {
	JobRPCServer
	PipelineRPCServer
}

func NewQueryTaskRequest() *QueryJobTaskRequest {
	return &QueryJobTaskRequest{
		Page: request.NewDefaultPageRequest(),
		Ids:  []string{},
	}
}

func NewRunJobRequest() *RunJobRequest {
	return &RunJobRequest{}
}

func NewRunTaskRequest(spec string, params *job.VersionedRunParam) *RunTaskRequest {
	return &RunTaskRequest{
		JobSpec: spec,
		Params:  params,
	}
}
