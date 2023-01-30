package runner

import (
	"context"

	"github.com/infraboard/mpaas/apps/task"
)

type Runner interface {
	Run(context.Context, *RunRequest) (*task.Task, error)
}

type RunRequest struct {
	RunnerSpec string
}
