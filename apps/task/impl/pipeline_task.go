package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/pipeline"
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

// Pipeline中任务有变化时,
// 如果执行成功则 继续执行, 如果失败则标记Pipeline结束
// 当所有任务成功结束时标记Pipeline执行成功
func (i *impl) PipelineTaskStatusChanged(ctx context.Context, in *task.JobTask) (
	*task.PipelineTaskSet, error) {
	if in == nil && in.Status == nil {
		return nil, exception.NewBadRequest("job task or job task status is nil")
	}

	if in.Spec.Pipeline == "" {
		return nil, exception.NewBadRequest("Pipeline Id参数缺失")
	}

	// 获取Pipeline
	p, err := i.pipeline.DescribePipeline(ctx, pipeline.NewDescribePipelineRequest(in.Spec.Pipeline))
	if err != nil {
		return nil, err
	}

	// 判断Task状态

	i.log.Debug(p)

	return nil, nil
}

// 查询Pipeline任务
func (i *impl) QueryPipelineTask(ctx context.Context, in *task.QueryPipelineTaskRequest) (
	*task.PipelineTaskSet, error) {
	return nil, nil
}
