package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/runner"
)

func (i *impl) RunJob(ctx context.Context, in *pipeline.RunJobRequest) (
	*task.JobTask, error) {
	ins := task.NewJobTask(in)

	// 1. 查询需要执行的Job
	req := job.NewDescribeJobRequest(in.Job)
	j, err := i.job.DescribeJob(ctx, req)
	if err != nil {
		return nil, err
	}
	ins.Job = j

	// 2. 执行Job
	r := runner.GetRunner(j.Spec.RunnerType)
	status, err := r.Run(ctx, task.NewRunTaskRequest(j.Spec.RunnerSpec, in.Params))
	if err != nil {
		return nil, err
	}
	ins.Status = status

	// 3. 保存任务
	if _, err := i.jcol.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a task document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryJobTask(ctx context.Context, in *task.QueryJobTaskRequest) (
	*task.JobTaskSet, error) {
	r := newQueryRequest(in)
	resp, err := i.jcol.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := task.NewTaskSet()
	// 循环
	for resp.Next(ctx) {
		ins := task.NewDefaultTask()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode deploy error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.jcol.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (i *impl) UpdateJobTaskStatus(ctx context.Context, in *task.UpdateJobTaskStatusRequest) (
	*task.JobTask, error) {
	return nil, nil
}

// 任务执行详情
func (i *impl) DescribeJobTask(ctx context.Context, in *task.DescribeJobTaskRequest) (
	*task.JobTask, error) {
	return nil, nil
}

// 删除任务
func (i *impl) DeleteJobTask(ctx context.Context, in *task.DeleteJobTaskRequest) (
	*task.JobTaskSet, error) {
	return nil, nil
}
