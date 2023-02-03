package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/task"
)

// 执行Pipeline
func (i *impl) RunPipeline(ctx context.Context, in *task.RunPipelineRequest) (
	*task.PipelineTask, error) {
	// 1. 查询需要执行的Pipeline
	p, err := i.pipeline.DescribePipeline(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(p)

	// 从pipeline 取出需要执行的任务
	p.GetFirstJob()

	return nil, nil
}

// 查询Pipeline任务
func (i *impl) QueryPipelineTask(ctx context.Context, in *task.QueryPipelineTaskRequest) (
	*task.PipelineTaskSet, error) {
	return nil, nil
}
