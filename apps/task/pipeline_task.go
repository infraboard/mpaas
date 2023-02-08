package task

import (
	"time"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/rs/xid"
)

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

// 返回下个需要执行的JobTask
func (p *PipelineTask) NextRun() *JobTask {
	return nil
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
		Name:     s.Name,
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

func NewDescribePipelineTaskRequest(id string) *DescribePipelineTaskRequest {
	return &DescribePipelineTaskRequest{
		Id: id,
	}
}
