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
	i.log.Debug(j)

	r := runner.GetRunner(j.Spec.RunnerType)
	t, err := r.Run(ctx, in)
	if err != nil {
		return nil, err
	}

	return t, nil
}
