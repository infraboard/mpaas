package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 执行Pipeline
func (i *impl) RunPipeline(ctx context.Context, in *task.RunPipelineRequest) (
	*task.PipelineTask, error) {
	// 检查请求
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 查询需要执行的Pipeline
	p, err := i.pipeline.DescribePipeline(ctx, pipeline.NewDescribePipelineRequest(in.PipelineId))
	if err != nil {
		return nil, err
	}

	// 检查Pipeline状态
	if err := i.CheckPipelineAllowRun(ctx, p); err != nil {
		return nil, err
	}

	// 从pipeline 取出需要执行的任务
	ins := task.NewPipelineTask(p)
	t := ins.GetFirstJobTask()
	if t == nil {
		return nil, fmt.Errorf("not job task to run")
	}

	// 保存Job Task, 所有JobTask 批量生成, 全部处于Pendding状态, 然后入库 等待状态更新
	err = i.JobTaskBatchSave(ctx, ins.JobTasks())
	if err != nil {
		return nil, err
	}

	// 运行Pipeline的一些准备工作
	err = i.PreparePipelineTask(ctx, ins)
	if err != nil {
		return nil, err
	}
	defer i.CleanPipelinTask(ctx, ins)

	// 运行 第一个Job, 驱动Pipeline执行
	ins.MarkedRunning()
	resp, err := i.RunJob(ctx, t.Spec)
	if err != nil {
		return nil, err
	}
	t.Update(resp.Job, resp.Status)

	// 保存Pipeline状态
	if _, err := i.pcol.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a pipeline task document error, %s", err)
	}

	return ins, nil
}

// 为Pipeline Task创建一个volume(通过创建secret 挂入), 用于记录Task运行时 输出的一些中间变量
func (i *impl) PreparePipelineTask(ctx context.Context, in *task.PipelineTask) error {
	return nil
}

// Pipeline Task运行结束时的一些清理工作, 比如删除中间生成的共享卷
func (i *impl) CleanPipelinTask(ctx context.Context, in *task.PipelineTask) {
}

func (i *impl) CheckPipelineAllowRun(ctx context.Context, ins *pipeline.Pipeline) error {
	// 1. 检查审核状态
	if ins.Spec.ApprovalId != "" {
		a, err := i.approval.DescribeApproval(
			ctx, approval.NewDescribeApprovalRequest(ins.Spec.ApprovalId),
		)
		if err != nil {
			return err
		}

		if !a.Status.IsAllowPublish() {
			return fmt.Errorf("当前状态: %s 不允许发布", a.Status.Stage)
		}
	}

	// 2. 检查当前pipeline是否已经处于运行中
	if !ins.Spec.IsParallel {
		// 查询当前pipeline最新的任务状态
		req := task.NewQueryPipelineTaskRequest()
		req.PipelineId = ins.Meta.Id
		req.Page.PageSize = 1
		set, err := i.QueryPipelineTask(ctx, req)
		if err != nil {
			return err
		}
		// 没有最近的任务
		if set.Len() == 0 {
			return nil
		}

		if set.Items[0].IsActive() {
			return fmt.Errorf("流水线当前处于运行中, 运行完成后才能运行")
		}
	}

	return nil
}

// Pipeline中任务有变化时,
// 如果执行成功则 继续执行, 如果失败则标记Pipeline结束
// 当所有任务成功结束时标记Pipeline执行成功
func (i *impl) PipelineTaskStatusChanged(ctx context.Context, in *task.JobTask) (
	*task.PipelineTask, error) {
	if in == nil && in.Status == nil {
		return nil, exception.NewBadRequest("job task or job task status is nil")
	}

	if in.Spec.PipelineTask == "" {
		return nil, exception.NewBadRequest("Pipeline Id参数缺失")
	}

	// 获取Pipeline Task, 因为Job Task是先保存在触发的回调, 这里获取的Pipeline Task是最新的
	descReq := task.NewDescribePipelineTaskRequest(in.Spec.PipelineTask)
	p, err := i.DescribePipelineTask(ctx, descReq)
	if err != nil {
		return nil, err
	}

	switch in.Status.Stage {
	case task.STAGE_PENDDING, task.STAGE_ACTIVE, task.STAGE_CANCELING:
		// Pipeline Task状态无变化
		return p, nil
	case task.STAGE_CANCELED:
		// 任务取消, pipeline 取消执行
		p.MarkedCanceled()
		return i.updatePipelineStatus(ctx, p)
	case task.STAGE_FAILED:
		// 任务执行失败, 更新Pipeline状态为失败
		if !in.Spec.IgnoreFailed {
			p.MarkedFailed()
			return i.updatePipelineStatus(ctx, p)
		}
	case task.STAGE_SUCCEEDED:
		// 任务运行成功, pipeline继续执行
	}

	// task执行成功或者忽略执行失败, 此时pipeline 仍然处于运行中, 需要获取下一个任务执行
	nexts, err := p.NextRun()
	if err != nil {
		return nil, err
	}

	// 如果没有需要执行的任务, Pipeline执行结束, 更新Pipeline状态为成功
	if nexts == nil && nexts.Len() == 0 {
		p.MarkedSuccess()
		return i.updatePipelineStatus(ctx, p)
	}

	// 如果有需要执行的JobTask, 继续执行
	for index := range nexts.Items {
		item := nexts.Items[index]
		t, err := i.RunJob(ctx, item.Spec)
		if err != nil {
			return nil, err
		}
		item.Status = t.Status
		item.Job = t.Job
	}

	return p, nil
}

// 更新Pipeline状态
func (i *impl) updatePipelineStatus(ctx context.Context, in *task.PipelineTask) (*task.PipelineTask, error) {
	// pipeline 状态更新回调
	i.updateCallback(ctx, in)

	in.Meta.UpdateAt = time.Now().Unix()
	if _, err := i.pcol.UpdateByID(ctx, in.Meta.Id, bson.M{"$set": bson.M{"status": in.Status}}); err != nil {
		return nil, exception.NewInternalServerError("update task(%s) document error, %s",
			in.Meta.Id, err)
	}
	return in, nil
}

func (i *impl) updateCallback(ctx context.Context, in *task.PipelineTask) {
	// pipeline task执行结束, 更新发布状态
	approvalId := in.Pipeline.Spec.ApprovalId
	if approvalId != "" && in.IsComplete() {
		req := approval.NewUpdateApprovalStatusRequest(approvalId)
		switch in.Status.Stage {
		case task.STAGE_FAILED:
			req.Status.Update(approval.STAGE_FAILED)
		case task.STAGE_SUCCEEDED:
			req.Status.Update(approval.STAGE_SUCCEEDED)
		case task.STAGE_CANCELED:
			req.Status.Update(approval.STAGE_CANCELED)
		}
		_, err := i.approval.UpdateApprovalStatus(ctx, req)
		if err != nil {
			i.log.Error(err)
		}
	}
}

// 查询Pipeline任务
func (i *impl) QueryPipelineTask(ctx context.Context, in *task.QueryPipelineTaskRequest) (
	*task.PipelineTaskSet, error) {
	r := newQueryPipelineTaskRequest(in)
	resp, err := i.pcol.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find pipeline task error, error is %s", err)
	}

	set := task.NewPipelineTaskSet()
	// 循环
	for resp.Next(ctx) {
		ins := task.NewDefaultPipelineTask()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode pipeline task  error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.pcol.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get pipeline task count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 查询Pipeline任务详情
func (i *impl) DescribePipelineTask(ctx context.Context, in *task.DescribePipelineTaskRequest) (
	*task.PipelineTask, error) {
	filter := bson.M{"_id": in.Id}

	ins := task.NewDefaultPipelineTask()
	if err := i.pcol.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("pipeline task %s not found", in.Id)
		}

		return nil, exception.NewInternalServerError("find pipeline task %s error, %s", in.Id, err)
	}

	// 补充该PipelineTask管理的JobTask
	query := task.NewQueryTaskRequest()
	query.PipelineTaskId = in.Id
	tasks, err := i.QueryJobTask(ctx, query)
	if err != nil {
		return nil, err
	}

	// 将tasks 填充给pipeline task
	for i := range tasks.Items {
		t := tasks.Items[i]
		stage := ins.GetStage(t.Spec.StageName)
		if stage != nil {
			stage.Add(t)
		}
	}

	return ins, nil
}

// 删除Pipeline任务详情
func (i *impl) DeletePipelineTask(ctx context.Context, in *task.DeletePipelineTaskRequest) (
	*task.PipelineTask, error) {
	ins, err := i.DescribePipelineTask(ctx, task.NewDescribePipelineTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}
	// TODO: 运行中的流水线不运行删除, 先取消 才能删除

	// 删除该Pipeline下所有的Job Task
	tasks := ins.Status.JobTasks()
	for index := range tasks.Items {
		t := tasks.Items[index]

		// 没有运行过的任务不需要清理
		if t.Status.Stage.Equal(task.STAGE_PENDDING) {
			continue
		}

		_, err := i.DeleteJobTask(ctx, task.NewDeleteJobTaskRequest(t.Spec.Id))
		if err != nil {
			if !exception.IsNotFoundError(err) {
				return nil, err
			}
		}
	}

	if err := i.deletecluster(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
