package task

import "github.com/infraboard/mpaas/apps/job"

func NewTask(req *RunJobRequest, job *job.Job, status *Status) *Task {
	return &Task{
		Spec:   req,
		Job:    job,
		Status: status,
	}
}

func NewStatus() *Status {
	return &Status{}
}
