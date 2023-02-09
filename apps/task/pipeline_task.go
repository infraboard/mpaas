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
	t := NewDefaultPipelineTask()
	t.Pipeline = p
	return t
}

func NewDefaultPipelineTask() *PipelineTask {
	return &PipelineTask{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
		Status:   NewPipelineTaskStatus(),
	}
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

// 返回下个需要执行的JobTask, 允许一次并行执行多个(批量执行)
func (p *PipelineTask) NextRun() *JobTaskSet {
	return nil
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

func (s *PipelineTaskStatus) AddStage(item *StageStatus) {
	s.StageStatus = append(s.StageStatus, item)
}

func NewStageStatus(s *pipeline.Stage) *StageStatus {
	status := &StageStatus{
		Spec:     s,
		JobTasks: []*JobTask{},
	}

	for i := range s.Jobs {
		req := s.Jobs[i]
		jobTask := NewJobTask(req)
		status.Add(jobTask)
	}

	return status
}

func (s *StageStatus) Add(item *JobTask) {
	s.JobTasks = append(s.JobTasks, item)
}

func (s *StageStatus) NextRun() []*JobTask {
	// 并行任务 返回该Stage所有等待执行的job
	if s.Spec.IsParallel {
		return s.JobTasks
	}

	return nil
}

func NewDescribePipelineTaskRequest(id string) *DescribePipelineTaskRequest {
	return &DescribePipelineTaskRequest{
		Id: id,
	}
}
