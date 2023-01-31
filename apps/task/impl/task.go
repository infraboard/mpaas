package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/runner"
)

func (i *impl) RunJob(ctx context.Context, in *task.RunJobRequest) (
	*task.Task, error) {
	// 1. 查询需要执行的Job
	j, err := i.job.DescribeJob(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 2. 执行Job
	r := runner.GetRunner(j.Spec.RunnerType)
	t, err := r.Run(ctx, task.NewRunTaskRequest(j.Spec.RunnerSpec, in.Params))
	if err != nil {
		return nil, err
	}

	return t, nil
}
