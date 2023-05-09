package pipeline

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/pb/resource"
	job "github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/common/format"
	"github.com/rs/xid"
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

	req.BuildNumber()

	d := &Pipeline{
		Meta: resource.NewMeta(),
		Spec: req,
	}
	return d, nil
}

// 注入编号
func (req *CreatePipelineRequest) BuildNumber() {
	for m := range req.Stages {
		stage := req.Stages[m]
		stage.Number = int32(m) + 1
		for n := range stage.Jobs {
			j := stage.Jobs[n]
			j.Number = int32(n) + 1
		}
	}
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
		*resource.Meta
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
		With:         []*job.RunParam{},
		Stages:       []*Stage{},
		Webhooks:     []*WebHook{},
		MentionUsers: []*MentionUser{},
		Labels:       map[string]string{},
	}
}

func (req *CreatePipelineRequest) ToYAML() string {
	yml, err := yaml.Marshal(req)
	if err != nil {
		panic(err)
	}
	return string(yml)
}

func (req *CreatePipelineRequest) MatchedWebHooks(event string) []*WebHook {
	hooks := []*WebHook{}
	for i := range req.Webhooks {
		h := req.Webhooks[i]
		if h.IsMatch(event) {
			hooks = append(hooks, h)
		}
	}

	return hooks
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

func (r *RunJobRequest) VersionName(version string) string {
	if strings.Contains(r.JobName, job.UNIQ_VERSION_SPLITER) {
		return r.JobName
	}
	return fmt.Sprintf("%s%s%s", r.JobName, job.UNIQ_VERSION_SPLITER, version)
}

func (r *RunJobRequest) Enabled() bool {
	return !r.SkipRun
}

func (r *RunJobRequest) MatchedWebHooks(event string) []*WebHook {
	hooks := []*WebHook{}
	for i := range r.Webhooks {
		h := r.Webhooks[i]
		if h.IsMatch(event) {
			hooks = append(hooks, h)
		}
	}

	return hooks
}

func (r *RunJobRequest) AddWebhook(items ...*WebHook) {
	r.Webhooks = append(r.Webhooks, items...)
}

func (r *RunJobRequest) AddMentionUser(items ...*MentionUser) {
	r.MentionUsers = append(r.MentionUsers, items...)
}

func (r *RunJobRequest) GetRunParamsVersion() string {
	if r.RunParams != nil {
		return r.RunParams.Version
	}

	return ""
}

func (r *RunJobRequest) BuildSearchLabel() {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	if r.RunParams == nil {
		return
	}

	lables := r.RunParams.SearchLabels()
	for k, v := range lables {
		r.Labels[k] = v
	}
}

func (r *RunJobRequest) SetDefault() {
	if r.TaskId == "" {
		r.TaskId = "task-" + xid.New().String()
	}
	if r.UpdateToken == "" {
		r.UpdateToken = xid.New().String()
	}
	if r.RunParams == nil {
		r.RunParams = job.NewVersionedRunParam("")
	}
	if r.RollbackParams == nil {
		r.RollbackParams = job.NewVersionedRunParam("")
	}
	if r.Webhooks == nil {
		r.Webhooks = []*WebHook{}
	}
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
}

func (r *RunJobRequest) GetJobShortName() string {
	nl := strings.Split(r.JobName, "@")
	if len(nl) > 0 && nl[0] != "" {
		return nl[0]
	}

	return r.JobName
}

func NewQueryPipelineRequestFromHTTP(r *http.Request) *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: request.NewPageRequestFromHTTP(r),
	}
}

func NewQueryPipelineRequest() *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewRunPipelineRequest(pipelineId string) *RunPipelineRequest {
	return &RunPipelineRequest{
		PipelineId: pipelineId,
		RunParams:  []*job.RunParam{},
		Labels:     make(map[string]string),
	}
}

func (req *RunPipelineRequest) AddRunParam(params ...*job.RunParam) {
	req.RunParams = append(req.RunParams, params...)
}

func (req *RunPipelineRequest) Validate() error {
	return validate.Validate(req)
}

func (req *RunPipelineRequest) ToJson() string {
	return format.Prettify(req)
}

func NewWebHook(url string) *WebHook {
	return &WebHook{
		Url:    url,
		Header: map[string]string{},
		Events: []string{},
	}
}

func (h *WebHook) IsMatch(event string) bool {
	if len(h.Events) == 0 {
		return true
	}

	for _, e := range h.Events {
		if strings.EqualFold(e, event) {
			return true
		}
	}

	return false
}

// 显示名称
func (h *WebHook) ShowName() string {
	return fmt.Sprintf("%s[%s]", h.Description, h.Url)
}

func NewMentionUser(username string) *MentionUser {
	return &MentionUser{
		UserName:    username,
		Events:      []string{},
		NotifyTypes: []notify.NOTIFY_TYPE{},
	}
}

func (m *MentionUser) AddEvent(events ...string) {
	m.Events = append(m.Events, events...)
}

func (m *MentionUser) AddNotifyType(nts ...notify.NOTIFY_TYPE) {
	m.NotifyTypes = append(m.NotifyTypes, nts...)
}

func (m *MentionUser) IsMatch(event string) bool {
	if len(m.Events) == 0 {
		return true
	}

	for _, e := range m.Events {
		if strings.EqualFold(e, event) {
			return true
		}
	}

	return false
}
