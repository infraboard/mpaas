package task

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/rs/xid"
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
func (s *JobTaskSet) GetJobTaskByStage(stage STAGE) (jobs []*JobTask) {
	for i := range s.Items {
		item := s.Items[i]
		if item.Status != nil && item.Status.Stage.Equal(stage) {
			jobs = append(jobs, item)
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
	if req.Id == "" {
		req.Id = xid.New().String()
	}

	return &JobTask{
		Meta:   NewMeta(),
		Spec:   req,
		Job:    nil,
		Status: NewJobTaskStatus(),
	}
}

func (p *JobTask) SystemVariable() (items []*job.RunParam) {
	items = append(items,
		job.NewRunParam(
			job.SYSTEM_VARIABLE_PIPELINE_TASK_ID,
			p.Spec.PipelineTask,
		),
		job.NewRunParam(
			job.SYSTEM_VARIABLE_JOB_TASK_ID,
			p.Spec.Id,
		),
	)
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

func (s *JobTask) ShowTitle() string {
	return fmt.Sprintf("任务[%s]当前状态: %s", s.Spec.JobName, s.Status.Stage.String())
}

func NewJobTaskStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt:            time.Now().Unix(),
		TemporaryResources: []*TemporaryResource{},
	}
}

func (t *JobTaskStatus) MarkedRunning() {
	t.StartAt = time.Now().Unix()
	t.Stage = STAGE_ACTIVE
}

func (t *JobTaskStatus) MarkedSuccess() {
	t.Stage = STAGE_SUCCEEDED
	t.EndAt = time.Now().Unix()
}

func (t *JobTaskStatus) IsComplete() bool {
	return t.Stage > 10
}

func (t *JobTaskStatus) Update(req *UpdateJobTaskStatusRequest) {
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

func (t *JobTaskStatus) AddTemporaryResource(items ...*TemporaryResource) {
	t.TemporaryResources = append(t.TemporaryResources, items...)
}

func NewTemporaryResource(kind, name string) *TemporaryResource {
	return &TemporaryResource{
		Kind:     kind,
		Name:     name,
		CreateAt: time.Now().Unix(),
	}
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
		kvs := strings.Split(l, "=")
		if len(kvs) != 2 {
			return nil, fmt.Errorf("环境变量格式错误: %s", l)
		}
		k, v := kvs[0], kvs[1]
		env := NewRuntimeEnv(strings.TrimSpace(k), strings.TrimSpace(v))
		envs = append(envs, env)
	}

	return envs, nil
}

func (r *RuntimeEnv) FileLine() (line []byte) {
	return []byte(fmt.Sprintf("%s=%s\n", r.Name, r.Value))
}
