package task

import (
	"time"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/rs/xid"
)

func NewPipelineTaskSet() *PipelineTaskSet {
	return &PipelineTaskSet{
		Items: []*PipelineTask{},
	}
}

func (s *PipelineTaskSet) Add(item *PipelineTask) {
	s.Items = append(s.Items, item)
}

func NewPipelineTask(p *pipeline.Pipeline) *PipelineTask {
	pt := NewDefaultPipelineTask()
	pt.Pipeline = p

	// 初始化所有的JobTask
	for i := range p.Spec.Stages {
		spec := p.Spec.Stages[i]
		ss := NewStageStatus(spec, pt.Id)
		pt.Status.AddStage(ss)
	}
	return pt
}

func NewDefaultPipelineTask() *PipelineTask {
	return &PipelineTask{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
		Status:   NewPipelineTaskStatus(),
	}
}

func (p *PipelineTask) MarkRunning() {
	p.Status.Stage = STAGE_ACTIVE
	p.Status.StartAt = time.Now().Unix()
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
	if p.Status != nil {
		return set
	}

	return p.Status.JobTasks()
}

func (p *PipelineTask) GetStageStatusByName(name string) *StageStatus {
	if p.Status != nil {
		return p.Status.GetStageStatusByName(name)
	}

	return nil
}

// 返回下个需要执行的JobTask, 允许一次并行执行多个(批量执行)
func (p *PipelineTask) NextRun() *JobTaskSet {
	return p.Status.NextRun()
}

func (p *PipelineTask) GetJobTask(id string) *JobTask {
	return p.Status.GetJobTask(id)
}

// 返回下个需要执行的JobTask
func (p *PipelineTask) MarkSuccess() {
	p.Status.Stage = STAGE_SUCCEEDED
	p.Status.EndAt = time.Now().Unix()
}

// 返回下个需要执行的JobTask
func (p *PipelineTask) MarkFailed() {
	p.Status.Stage = STAGE_FAILED
	p.Status.EndAt = time.Now().Unix()
}

func NewPipelineTaskStatus() *PipelineTaskStatus {
	return &PipelineTaskStatus{
		StageStatus: []*StageStatus{},
	}
}

func (p *PipelineTaskStatus) GetStageStatusByName(name string) *StageStatus {
	for i := range p.StageStatus {
		stage := p.StageStatus[i]
		if stage.Spec.Name == name {
			return stage
		}
	}

	return nil
}

func (p *PipelineTaskStatus) JobTasks() *JobTaskSet {
	set := NewJobTaskSet()
	for i := range p.StageStatus {
		stage := p.StageStatus[i]
		set.Add(stage.JobTasks...)
	}
	return set
}

func (s *PipelineTaskStatus) NextRun() *JobTaskSet {
	set := NewJobTaskSet()
	for i := range s.StageStatus {
		s := s.StageStatus[i]
		set.Add(s.NextRun()...)
	}
	return set
}

func (s *PipelineTaskStatus) AddStage(item *StageStatus) {
	s.StageStatus = append(s.StageStatus, item)
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
		Spec:     s,
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

func (s *StageStatus) Add(item *JobTask) {
	s.JobTasks = append(s.JobTasks, item)
}

// 根据Job Task id获取当前stage中的Job Task
func (s *StageStatus) GetJobTask(id string) *JobTask {
	for i := range s.JobTasks {
		item := s.JobTasks[i]
		if item.Id == id {
			return item
		}
	}
	return nil
}

func (s *StageStatus) NextRun() (jobs []*JobTask) {
	tasks := s.PenddingJobs()
	// 并行任务 返回该Stage所有等待执行的job
	if s.Spec.IsParallel {
		jobs = append(jobs, tasks...)
		return
	}

	// 串行任务取第一个
	if len(tasks) > 0 {
		jobs = append(jobs, tasks[0])
		return
	}

	return
}

func (s *StageStatus) PenddingJobs() (jobs []*JobTask) {
	for i := range s.JobTasks {
		t := s.JobTasks[i]
		if t.Status.Stage.Equal(STAGE_PENDDING) {
			jobs = append(jobs, t)
		}
	}
	return
}

func NewDescribePipelineTaskRequest(id string) *DescribePipelineTaskRequest {
	return &DescribePipelineTaskRequest{
		Id: id,
	}
}
