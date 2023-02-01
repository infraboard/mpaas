package task

import (
	"time"

	"github.com/infraboard/mpaas/apps/job"
)

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

func (s *TaskSet) Add(task *Task) {
	s.Items = append(s.Items, task)
}

func NewDefaultTask() *Task {
	return NewTask(NewRunJobRequest(), nil, NewStatus())
}

func NewTask(req *RunJobRequest, job *job.Job, status *Status) *Task {
	return &Task{
		Spec:   req,
		Job:    job,
		Status: status,
	}
}

func NewStatus() *Status {
	return &Status{
		StartAt: time.Now().Unix(),
	}
}
