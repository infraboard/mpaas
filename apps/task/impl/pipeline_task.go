package impl

import (
	"context"

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
	ins := task.NewPipelineTask(p)

	// 从pipeline 取出需要执行的任务
	jt := ins.GetFirstJobTask()

	// 运行Job
	resp, err := i.RunJob(ctx, jt.Spec)
	if err != nil {
		return nil, err
	}
	jt.Update(resp.Job, resp.Status)

	return ins, nil
}

// 查询Pipeline任务
func (i *impl) QueryPipelineTask(ctx context.Context, in *task.QueryPipelineTaskRequest) (
	*task.PipelineTaskSet, error) {
	return nil, nil
}
