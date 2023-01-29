package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/task"
)

func (i *impl) RunJob(ctx context.Context, in *task.RunJobRequest) (
	*task.Task, error) {
	// 1. 查询需要执行的Job
	i.job.QueryJob(ctx, nil)

	return nil, nil
}
