package task

import job "github.com/infraboard/mpaas/apps/job"

const (
	AppName = "tasks"
)

type Service interface {
	RPCServer
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
