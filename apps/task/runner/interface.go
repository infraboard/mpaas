package runner

import (
	"context"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
)

var (
	runners = map[job.RUNNER_TYPE]Register{}
)

type Register interface {
	RunnerType() job.RUNNER_TYPE
	Init() error
	Runner
}

type Runner interface {
	Run(context.Context, *task.RunTaskRequest) (*task.Task, error)
}

func Init() error {
	for i := range runners {
		r := runners[i]
		if err := r.Init(); err != nil {
			return err
		}
	}
	return nil
}

func Registry(runner Register) {
	runners[runner.RunnerType()] = runner
}

func GetRunner(rt job.RUNNER_TYPE) Runner {
	return runners[rt]
}
