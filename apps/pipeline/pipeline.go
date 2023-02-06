package pipeline

import (
	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
)

func NewRunJobRequest(jobName string) *RunJobRequest {
	return &RunJobRequest{
		Job:    jobName,
		Params: job.NewVersionedRunParam(""),
	}
}

func NewQueryPipelineRequest() *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
