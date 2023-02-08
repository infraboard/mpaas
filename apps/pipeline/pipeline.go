package pipeline

import (
	"fmt"
	"time"

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

func (h *WebHook) StartSend() {
	if h.Status == nil {
		h.Status = &WebHookStatus{}
	}
	h.Status.StartAt = time.Now().UnixMilli()
}

func (h *WebHook) SendFailed(format string, a ...interface{}) {
	if h.Status.StartAt != 0 {
		h.Status.Cost = time.Now().UnixMilli() - h.Status.StartAt
	}
	h.Status.Message = fmt.Sprintf(format, a...)
}

func (h *WebHook) Success(message string) {
	if h.Status.StartAt != 0 {
		h.Status.Cost = time.Now().UnixMilli() - h.Status.StartAt
	}
	h.Status.Success = true
	h.Status.Message = message
}

func (h *WebHook) IsMatch(t string) bool {
	for _, e := range h.Events {
		if e == t {
			return true
		}
	}

	return false
}
