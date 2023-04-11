package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"
	"unicode"

	"github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/common/format"
)

func NewJobTaskSet() *JobTaskSet {
	return &JobTaskSet{
		Items: []*JobTask{},
	}
}

func (s *JobTaskSet) Add(tasks ...*JobTask) {
	s.Items = append(s.Items, tasks...)
}

func (s *JobTaskSet) Len() int {
	return len(s.Items)
}

func (s *JobTaskSet) ToDocs() (docs []any) {
	for i := range s.Items {
		docs = append(docs, s.Items[i])
	}
	return
}

func (s *JobTaskSet) HasStage(stage STAGE) bool {
	for i := range s.Items {
		item := s.Items[i]
		if item.Status != nil && item.Status.Stage.Equal(stage) {
			return true
		}
	}
	return false
}

// 查询Stage中 等待执行的Job Task
func (s *JobTaskSet) GetJobTaskByStage(stage STAGE) (tasks []*JobTask) {
	for i := range s.Items {
		item := s.Items[i]
		if item.Status != nil && item.Status.Stage.Equal(stage) {
			tasks = append(tasks, item)
		}
	}
	return
}

func NewDefaultJobTask() *JobTask {
	req := pipeline.NewRunJobRequest("")
	return NewJobTask(req)
}

func NewMeta() *Meta {
	return &Meta{
		CreateAt: time.Now().Unix(),
	}
}

func NewJobTask(req *pipeline.RunJobRequest) *JobTask {
	req.SetDefault()
	t := &JobTask{
		Meta:   NewMeta(),
		Spec:   req,
		Job:    nil,
		Status: NewJobTaskStatus(),
	}

	if t.Spec.SkipRun {
		t.Status.Stage = STAGE_SKIPPED
		t.Status.Message = "skip run"
	}
	return t
}

var (
	// 关于Go模版语法可以参考: https://www.tizi365.com/archives/85.html
	JOB_TASK_MARKDOWN_TEMPLATE = `
**开始时间: **
{{ .Status.StartAtFormat }}
**结束时间: **
{{ .Status.EndAtAtFormat }}
**任务参数: **
{{ range .Spec.RunParams.Params -}}
▫ *{{.Name}}:  {{.Value}}*
{{end}}
`
)

func (p *JobTask) MarkdownContent() string {
	buf := bytes.NewBuffer([]byte{})
	t := template.New("job task")
	tmpl, err := t.Parse(JOB_TASK_MARKDOWN_TEMPLATE)
	if err != nil {
		return err.Error()
	}

	err = tmpl.Execute(buf, p)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

var (
	// 关于Go模版语法可以参考: https://www.tizi365.com/archives/85.html
	JOB_TASK_HTML_TEMPLATE = `
开始时间: 
{{ .Status.StartAtFormat }}
结束时间: 
{{ .Status.EndAtAtFormat }}
任务参数: 
{{ range .Spec.RunParams.Params -}}
▫ *{{.Name}}:  {{.Value}}*
{{end}}
`
)

func (p *JobTask) HTMLContent() string {
	buf := bytes.NewBuffer([]byte{})
	t := template.New("job task")
	tmpl, err := t.Parse(JOB_TASK_HTML_TEMPLATE)
	if err != nil {
		return err.Error()
	}

	err = tmpl.Execute(buf, p)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func (p *JobTask) BuildSearchLabel() {
	if p.Job != nil && p.Job.Spec != nil {
		if p.Job.Spec.Labels == nil {
			p.Job.Spec.Labels = map[string]string{}
		}
		for k, v := range p.Job.Spec.Labels {
			p.Spec.Labels[k] = v
		}
	}

	p.Spec.BuildSearchLabel()
}

func (p *JobTask) GetVersionedRunParam() (*job.VersionedRunParam, error) {
	j := p.Job
	v := p.Spec.GetRunParamsVersion()
	params := j.GetVersionedRunParam(v)
	if params != nil {
		return params, nil
	}

	return nil, fmt.Errorf("job %s version: %s not found, allow version: %s",
		j.Spec.Name,
		v,
		j.AllowVersions(),
	)
}

func (p *JobTask) SystemRunParam() (items []*job.RunParam) {
	items = append(items,
		job.NewRunParam(
			job.SYSTEM_VARIABLE_JOB_TASK_ID,
			p.Spec.TaskId,
		).SetReadOnly(true).SetRequired(true),
	)

	if p.Spec.PipelineTask != "" {
		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_PIPELINE_TASK_ID,
				p.Spec.PipelineTask,
			).SetReadOnly(true).SetRequired(true),
		)
	}
	return
}

func (p *JobTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*pipeline.RunJobRequest
		*Meta
		*JobTaskStatus
		Job *job.Job `json:"job"`
	}{p.Spec, p.Meta, p.Status, p.Job})
}

func (t *JobTask) GetStatusDetail() string {
	if t.Status != nil {
		return t.Status.Detail
	}

	return ""
}

func (s *JobTask) HasJobSpec() bool {
	if s.Job != nil && s.Job.Spec != nil {
		return true
	}

	return false
}

func (t *JobTask) Update(job *job.Job, status *JobTaskStatus) {
	t.Job = job
	t.Status = status
}

func (t *JobTask) ToJson() string {
	return format.Prettify(t)
}

func (t *JobTask) ValidateToken(token string) error {
	if t.Spec.UpdateToken != token {
		return fmt.Errorf("update token invalidate")
	}
	return nil
}

func (s *JobTask) ShowTitle() string {
	return fmt.Sprintf("任务[%s]当前状态: %s", s.Spec.JobName, s.Status.Stage.String())
}

func NewJobTaskStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt:            time.Now().Unix(),
		TemporaryResources: []*TemporaryResource{},
	}
}

func (t *JobTaskStatus) MessageToError() error {
	return fmt.Errorf(t.Message)
}

func (t *JobTaskStatus) MarkedError(err error) {
	t.EndAt = time.Now().Unix()
	t.Stage = STAGE_FAILED
	t.Message = err.Error()
}

func (t *JobTaskStatus) MarkedRunning() {
	t.StartAt = time.Now().Unix()
	t.Stage = STAGE_ACTIVE
}

func (t *JobTaskStatus) MarkedSuccess() {
	t.Stage = STAGE_SUCCEEDED
	t.EndAt = time.Now().Unix()
}

func (t *JobTaskStatus) StartAtFormat() string {
	start := time.Unix(t.StartAt, 0)
	return start.Format("2006-01-02 03:04:05")
}

func (t *JobTaskStatus) EndAtAtFormat() string {
	start := time.Unix(t.EndAt, 0)
	return start.Format("2006-01-02 03:04:05")
}

func (t *JobTaskStatus) IsComplete() bool {
	return t.Stage >= STAGE_CANCELED
}

func (p *JobTaskStatus) AddWebhookStatus(items ...*CallbackStatus) {
	p.WebhookStatus = append(p.WebhookStatus, items...)
}

func (p *JobTaskStatus) AddNotifyStatus(items ...*CallbackStatus) {
	p.NotifyStatus = append(p.NotifyStatus, items...)
}

func (t *JobTaskStatus) UpdateStatus(req *UpdateJobTaskStatusRequest) {
	t.Stage = req.Stage
	t.Message = req.Message

	// 取消的任务 不需要更新detail详情
	if !t.Stage.Equal(STAGE_CANCELED) {
		t.Detail = req.Detail
	}

	// 如果没传递结束时间, 则自动生成结束时间
	if t.IsComplete() && t.EndAt == 0 {
		t.EndAt = time.Now().Unix()
	}
}

func (t *JobTaskStatus) UpdateOutput(req *UpdateJobTaskOutputRequest) {
	t.RuntimeEnvs = req.RuntimeEnvs
	if req.MarkdownOutput != "" {
		t.MarkdownOutput = req.MarkdownOutput
	}
}

func (t *JobTaskStatus) AddTemporaryResource(items ...*TemporaryResource) {
	t.TemporaryResources = append(t.TemporaryResources, items...)
}

func (t *JobTaskStatus) AddErrorEvent(format string, a ...any) {
	t.Events = append(t.Events, NewEvent(EVENT_LEVEL_ERROR, fmt.Sprintf(format, a...)))
}

func (t *JobTaskStatus) AddEvent(level EVENT_LEVEL, format string, a ...any) {
	t.Events = append(t.Events, NewEvent(level, fmt.Sprintf(format, a...)))
}

func (t *JobTaskStatus) GetTemporaryResource(kind, name string) *TemporaryResource {
	for i := range t.TemporaryResources {
		tr := t.TemporaryResources[i]
		if tr.Kind == kind && tr.Name == name {
			return tr
		}
	}
	return nil
}

func NewTemporaryResource(kind, name string) *TemporaryResource {
	return &TemporaryResource{
		Kind:     kind,
		Name:     name,
		CreateAt: time.Now().Unix(),
	}
}

func (r *TemporaryResource) IsReleased() bool {
	return r.ReleaseAt != 0
}

func ParseRuntimeEnvFromBytes(content []byte) ([]*RuntimeEnv, error) {
	envs := []*RuntimeEnv{}
	lines := []string{}
	line := []byte{}
	for _, c := range content {
		if c == '\n' {
			lines = append(lines, string(line))
			line = []byte{}
		} else {
			line = append(line, c)
		}
	}

	for _, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		kvs := strings.Split(l, "=")
		if len(kvs) != 2 {
			return nil, fmt.Errorf("环境变量格式错误: %s", kvs)
		}
		k, v := kvs[0], kvs[1]

		env := NewRuntimeEnv(k, strings.Trim(v, `"`))
		envs = append(envs, env)
	}

	return envs, nil
}

func (r *RuntimeEnv) FileLine() (line []byte) {
	return []byte(fmt.Sprintf("%s=%s\n", r.Name, r.Value))
}

func (r *RuntimeEnv) IsExport() bool {
	if r.Name == "" && unicode.IsUpper(rune(r.Name[0])) {
		return true
	}
	return false
}

func NewEvent(level EVENT_LEVEL, message string) *Event {
	return &Event{
		Time:    time.Now().Unix(),
		Level:   level,
		Message: message,
	}
}
