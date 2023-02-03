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

func (p *PipelineTask) GetFirstJobTask() *JobTask {
	for i := range p.Status.StageStatus {
		s := p.Status.StageStatus[i]
		if len(s.JobTasks) > 0 {
			return s.JobTasks[0]
		}
	}
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
		spec := s.Jobs[i]
		// 获取pipeline job参数, 构造task run 参数
		req := NewRunJobRequest(spec.JobName())
		jopParams := job.NewVersionedRunParam(spec.JobVersion())
		jopParams.Params = spec.RunParams()
		req.Params = jopParams

		jobTask := NewJobTask(req)
		status.Add(jobTask)
	}

	return status
}

func (s *StageStatus) Add(item *JobTask) {
	s.JobTasks = append(s.JobTasks, item)
}
