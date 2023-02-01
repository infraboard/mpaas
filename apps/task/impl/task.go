package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
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
	status, err := r.Run(ctx, task.NewRunTaskRequest(j.Spec.RunnerSpec, in.Params))
	if err != nil {
		return nil, err
	}

	// 3. 保存任务
	ins := task.NewTask(in, j, status)
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a task document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryTask(ctx context.Context, in *task.QueryTaskRequest) (
	*task.TaskSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

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
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
