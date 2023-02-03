package task

import (
	"time"

	"github.com/infraboard/mpaas/apps/job"
)

func NewTaskSet() *JobTaskSet {
	return &JobTaskSet{
		Items: []*JobTask{},
	}
}

func (s *JobTaskSet) Add(task *JobTask) {
	s.Items = append(s.Items, task)
}

func NewDefaultTask() *JobTask {
	return NewTask(NewRunJobRequest(), nil, NewStatus())
}

func NewTask(req *RunJobRequest, job *job.Job, status *JobTaskStatus) *JobTask {
	return &JobTask{
		Spec:   req,
		Job:    job,
		Status: status,
	}
}

func NewStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt: time.Now().Unix(),
	}
}
