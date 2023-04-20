package task

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/common/meta"
)

func NewPipelineTaskSet() *PipelineTaskSet {
	return &PipelineTaskSet{
		Items: []*PipelineTask{},
	}
}

func (s *PipelineTaskSet) Add(item *PipelineTask) {
	s.Items = append(s.Items, item)
}

func (s *PipelineTaskSet) Len() int {
	return len(s.Items)
}

func NewPipelineTask(p *pipeline.Pipeline, in *pipeline.RunPipelineRequest) *PipelineTask {
	pt := NewDefaultPipelineTask()
	pt.Pipeline = p
	pt.Params = in

	// 如果传入了id 则使用传入的id
	if in.PipelineTaskId != "" {
		pt.Meta.Id = in.PipelineTaskId
		in.PipelineTaskId = ""
	}

	// 初始化所有的JobTask
	for i := range p.Spec.Stages {
		spec := p.Spec.Stages[i]
		ss := NewStageStatus(spec, pt.Meta.Id)
		pt.Status.AddStage(ss)
	}
	return pt
}

func NewDefaultPipelineTask() *PipelineTask {
	return &PipelineTask{
		Meta:   meta.NewMeta(),
		Params: pipeline.NewRunPipelineRequest(""),
		Status: NewPipelineTaskStatus(),
	}
}

func (p *PipelineTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*pipeline.RunPipelineRequest
		*PipelineTaskStatus
		Pipeline *pipeline.Pipeline `json:"pipeline"`
	}{p.Meta, p.Params, p.Status, p.Pipeline})
}

func (p *PipelineTask) IsActive() bool {
	if p.Status != nil && p.Status.Stage.Equal(STAGE_ACTIVE) {
		return true
	}

	return false
}

func (p *PipelineTask) HasJobSpec() bool {
	if p.Pipeline != nil && p.Pipeline.Spec != nil {
		return true
	}

	return false
}

func (p *PipelineTask) MarkedRunning() {
	p.Status.Stage = STAGE_ACTIVE
	p.Status.StartAt = time.Now().Unix()
}

func (p *PipelineTask) MarkedSucceed() {
	p.Status.Stage = STAGE_SUCCEEDED
	p.Status.EndAt = time.Now().Unix()
}

func (p *PipelineTask) GetFirstJobTask() *JobTask {
	for i := range p.Status.StageStatus {
		s := p.Status.StageStatus[i]
		if len(s.JobTasks) > 0 {
			return s.JobTasks[0]
		}
	}
	return nil
}

func (p *PipelineTask) JobTasks() *JobTaskSet {
	set := NewJobTaskSet()
	if p.Status == nil {
		return set
	}

	return p.Status.JobTasks()
}

// 返回下个需要执行的JobTask, 允许一次并行执行多个(批量执行)
func (p *PipelineTask) NextRun() (*JobTaskSet, error) {
	set := NewJobTaskSet()
	var stage *StageStatus

	if p.Status == nil || p.Pipeline == nil {
		return set, nil
	}

	// 需要未执行完成的Job Tasks
	stages := p.Status.StageStatus
	for i := range stages {
		stage = stages[i]

		// 找出Stage中未执行完的Job Task
		set = stage.UnCompleteJobTask()

		// 如果找到 直接Break
		if set.Len() > 0 {
			break
		}
	}

	// 如果所有Stage寻找完，都没找到, 表示PipelineTask执行完成
	if set.Len() == 0 {
		return set, nil
	}

	// 如果这些未执行当中的Job Task 有处于运行中的, 不会执行下个一个任务
	if set.HasStage(STAGE_ACTIVE) {
		return set, exception.NewConflict("Stage 还处于运行中")
	}

	// 当未执行的任务中，没有运行中的时，剩下的就是需要被执行的任务
	tasks := set.GetJobTaskByStage(STAGE_PENDDING)

	nextTasks := NewJobTaskSet()
	stageSpec := p.Pipeline.GetStage(stage.Name)
	if stageSpec.IsParallel {
		// 并行任务 返回该Stage所有等待执行的job
		nextTasks.Add(tasks...)
	} else {
		// 串行任务取第一个
		nextTasks.Add(tasks[0])
	}

	return nextTasks, nil
}

func (p *PipelineTask) GetStage(name string) *StageStatus {
	if p.Status == nil || p.Pipeline == nil {
		return nil
	}

	stages := p.Status.StageStatus
	// 先查 StageStatus 是否有
	for i := range stages {
		stage := stages[i]
		if stage.Name == name {
			return stage
		}
	}

	// 如果没有动态创建
	stageSpec := p.Pipeline.GetStage(name)
	if stageSpec == nil {
		return nil
	}

	stage := NewStageStatus(stageSpec, p.Meta.Id)
	p.Status.AddStage(stage)

	return stage
}

func (p *PipelineTask) GetJobTask(id string) *JobTask {
	return p.Status.GetJobTask(id)
}

// Pipeline执行成功
func (p *PipelineTask) UUID() string {
	return fmt.Sprintf("jobtask-%s", p.Meta.Id)
}

// Pipeline执行成功
func (p *PipelineTask) MarkedSuccess() {
	p.Status.Stage = STAGE_SUCCEEDED
	p.Status.EndAt = time.Now().Unix()
}

// Pipeline执行失败
func (p *PipelineTask) MarkedFailed(err error) {
	p.Status.Stage = STAGE_FAILED
	p.Status.EndAt = time.Now().Unix()
	p.Status.Message = err.Error()
}

// Pipeline执行取消
func (p *PipelineTask) MarkedCanceled() {
	p.Status.Stage = STAGE_CANCELED
	p.Status.EndAt = time.Now().Unix()
}

// Pipeline执行取消
func (p *PipelineTask) IsComplete() bool {
	if p.Status != nil && p.Status.Stage >= STAGE_CANCELED {
		return true
	}

	return false
}

// Pipeline执行取消
func (p *PipelineTask) IsRunning() bool {
	if p.Status == nil {
		return false
	}

	return p.Status.Stage > STAGE_PENDDING && p.Status.Stage < STAGE_CANCELED
}

// 大写导出
func (s *PipelineTask) RuntimeRunParams() (envs []*job.RunParam) {
	if s.Status == nil {
		return
	}

	for i := range s.Status.RuntimeEnvs {
		env := s.Status.RuntimeEnvs[i]
		if env.IsExport() {
			envs = append(envs, job.NewRunParam(
				env.Name,
				env.Value,
			))
		}
	}

	return
}

func NewPipelineTaskStatus() *PipelineTaskStatus {
	return &PipelineTaskStatus{
		StageStatus: []*StageStatus{},
	}
}

func (p *PipelineTaskStatus) JobTasks() *JobTaskSet {
	set := NewJobTaskSet()
	for i := range p.StageStatus {
		stage := p.StageStatus[i]
		set.Add(stage.JobTasks...)
	}
	return set
}

func (s *PipelineTaskStatus) AddStage(item *StageStatus) {
	s.StageStatus = append(s.StageStatus, item)
}

func (s *PipelineTaskStatus) UpdateRuntimeEnv(updateBy string, envs []*RuntimeEnv) {
	for i := range envs {
		env := envs[i]

		// 未导出变量不更新
		if !env.IsExport() {
			continue
		}

		// 获取Pipeline中变量
		old := s.GetRuntimeEnv(env.Name)
		if old == nil {
			s.AddRuntimeEnv(env)
			return
		}

		// 不更新只读
		if old.ReadOnly {
			return
		}
		old.UpdateAt = time.Now().Unix()
		old.UpdateTaskId = updateBy
		old.Value = env.Value
	}
}

func (s *PipelineTaskStatus) AddRuntimeEnv(items ...*RuntimeEnv) {
	s.RuntimeEnvs = append(s.RuntimeEnvs, items...)
}

func (s *PipelineTaskStatus) GetRuntimeEnv(name string) *RuntimeEnv {
	for i := range s.RuntimeEnvs {
		env := s.RuntimeEnvs[i]
		if env.Name == name {
			return env
		}
	}
	return nil
}

func (s *PipelineTaskStatus) GetJobTask(id string) *JobTask {
	for i := range s.StageStatus {
		stage := s.StageStatus[i]
		jobTask := stage.GetJobTask(id)
		if jobTask != nil {
			return jobTask
		}
	}

	return nil
}

func NewStageStatus(s *pipeline.Stage, pipelineTaskId string) *StageStatus {
	status := &StageStatus{
		Name:     s.Name,
		JobTasks: []*JobTask{},
	}

	for i := range s.Jobs {
		req := s.Jobs[i]
		req.PipelineTask = pipelineTaskId
		req.StageName = s.Name
		jobTask := NewJobTask(req)
		status.Add(jobTask)
	}

	return status
}

// 获取Stage中未完成的任务, 包含 运行中和等待执行的
func (s *StageStatus) UnCompleteJobTask() *JobTaskSet {
	set := NewJobTaskSet()
	for i := range s.JobTasks {
		item := s.JobTasks[i]
		// stage中任何一个job task未完成, 该stage都处于未完成
		if item.Status != nil && !item.Status.IsComplete() {
			set.Add(item)
		}
	}

	return set
}

func (s *StageStatus) Add(item *JobTask) {
	s.JobTasks = append(s.JobTasks, item)
}

// 根据Job Task id获取当前stage中的Job Task
func (s *StageStatus) GetJobTask(id string) *JobTask {
	for i := range s.JobTasks {
		item := s.JobTasks[i]
		if item.Spec.TaskId == id {
			return item
		}
	}
	return nil
}

func NewReadOnlyRuntimeEnv(name, value string) *RuntimeEnv {
	env := NewRuntimeEnv(name, value)
	env.ReadOnly = true
	return env
}

func NewRuntimeEnv(name, value string) *RuntimeEnv {
	return &RuntimeEnv{
		Name:  name,
		Value: value,
	}
}
