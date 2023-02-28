package pipeline

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/http/request"
	job "github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/common/meta"
	"sigs.k8s.io/yaml"
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
		Meta: meta.NewMeta(),
		Spec: req,
	}
	return d, nil
}

func (p *Pipeline) Update(req *UpdatePipelineRequest) {
	p.Meta.UpdateAt = time.Now().Unix()
	p.Meta.UpdateBy = req.UpdateBy
	p.Spec = req.Spec
}

func (p *Pipeline) Patch(req *UpdatePipelineRequest) error {
	p.Meta.UpdateAt = time.Now().Unix()
	p.Meta.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(p.Spec, req.Spec)
}

func (p *Pipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreatePipelineRequest
	}{p.Meta, p.Spec})
}

func (p *Pipeline) GetStage(name string) *Stage {
	if p.Spec == nil {
		return nil
	}

	for i := range p.Spec.Stages {
		stage := p.Spec.Stages[i]
		if stage.Name == name {
			return stage
		}
	}

	return nil
}

func NewCreatePipelineRequestFromYAML(yml string) (*CreatePipelineRequest, error) {
	req := NewCreatePipelineRequest()

	err := yaml.Unmarshal([]byte(yml), req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func NewCreatePipelineRequest() *CreatePipelineRequest {
	return &CreatePipelineRequest{
		Stages: []*Stage{},
		Labels: map[string]string{},
	}
}

func (req *CreatePipelineRequest) ToYAML() string {
	yml, err := yaml.Marshal(req)
	if err != nil {
		panic(err)
	}
	return string(yml)
}

func (req *CreatePipelineRequest) AddStage(stages ...*Stage) {
	req.Stages = append(req.Stages, stages...)
}

func NewRunJobRequest(jobName string) *RunJobRequest {
	return &RunJobRequest{
		JobName:   jobName,
		RunParams: job.NewVersionedRunParam(""),
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
