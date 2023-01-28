package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/task"
)

func (i *impl) RunJob(ctx context.Context, in *task.RunJobRequest) (
	*task.Task, error) {
	return nil, nil
}
