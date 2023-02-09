package pipeline

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
	"github.com/rs/xid"
)

func NewPipelineSet() *PipelineSet {
	return &PipelineSet{
		Items: []*Pipeline{},
	}
}

func (s *PipelineSet) Add(item *Pipeline) {
	s.Items = append(s.Items, item)
}

func NewDefaultPipeline() *Pipeline {
	return &Pipeline{
		Spec: NewCreatePipelineRequest(),
	}
}

// New 新建一个部署配置
func New(req *CreatePipelineRequest) (*Pipeline, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Pipeline{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
		Spec:     req,
	}
	return d, nil
}

func NewCreatePipelineRequest() *CreatePipelineRequest {
	return &CreatePipelineRequest{
		Stages: []*Stage{},
		Labels: map[string]string{},
	}
}

func (req *CreatePipelineRequest) AddStage(stages ...*Stage) {
	req.Stages = append(req.Stages, stages...)
}

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
