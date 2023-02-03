package task

import (
	"time"

	job "github.com/infraboard/mpaas/apps/job"
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
		Name:   s.Name,
		Status: []*JobTask{},
	}

	for i := range s.Jobs {
		spec := s.Jobs[i]
		jopParams := job.NewVersionedRunParam(spec.JobVersion())
		req := NewRunJobRequest(spec.JobName(), jopParams)
		jobTask := NewJobTask(req)
		status.Add(jobTask)
	}

	return status
}

func (s *StageStatus) Add(item *JobTask) {
	s.Status = append(s.Status, item)
}
