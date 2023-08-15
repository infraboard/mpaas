package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/trigger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 执行Pipeline
func (i *impl) RunPipeline(ctx context.Context, in *pipeline.RunPipelineRequest) (
	*task.PipelineTask, error) {
	// 检查Pipeline请求参数
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 查询需要执行的Pipeline
	p, err := i.pipeline.DescribePipeline(ctx, pipeline.NewDescribePipelineRequest(in.PipelineId))
	if err != nil {
		return nil, err
	}

	// 检查Pipeline状态
	if err := i.CheckPipelineAllowRun(ctx, p, in.ApprovalId); err != nil {
		return nil, err
	}

	// 从pipeline 取出需要执行的任务
	ins := task.NewPipelineTask(p, in)
	t := ins.GetFirstJobTask()
	if t == nil {
		return nil, fmt.Errorf("not job task to run")
	}

	// dry run 不执行
	if in.DryRun {
		return ins, nil
	}

	// 保存Job Task, 所有JobTask 批量生成, 全部处于Pendding状态, 然后入库 等待状态更新
	err = i.JobTaskBatchSave(ctx, ins.JobTasks())
	if err != nil {
		return nil, err
	}

	// 保存Pipeline状态
	if _, err := i.pcol.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a pipeline task document error, %s", err)
	}

	// 运行 第一个Job, 驱动Pipeline执行
	ins.MarkedRunning()
	resp, err := i.RunJob(ctx, t.Spec)
	if err != nil {
		return nil, err
	}
	t.Update(resp.Job, resp.Status)

	// 更新状态
	ins, err = i.updatePipelineStatus(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) CheckPipelineAllowRun(ctx context.Context, ins *pipeline.Pipeline, approlvalId string) error {
	// 1. 检查审核状态
	if ins.Spec.RequiredApproval {
		if approlvalId == "" {
			return exception.NewBadRequest("流水线需要审核单才能运行")
		}
		a, err := i.approval.DescribeApproval(
			ctx, approval.NewDescribeApprovalRequest(approlvalId),
		)
		if err != nil {
			return err
		}

		if a.Pipeline.Meta.Id != ins.Meta.Id {
			return fmt.Errorf("审核单不属于当前执行的流水线")
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

		t := set.Items[0]
		if t.IsActive() {
			return fmt.Errorf("流水线【%s】当前处于运行中, 请取消或者等待之前运行结束", t.Pipeline.Spec.Name)
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

	runErrorJobTasks := []*task.UpdateJobTaskStatusRequest{}
	// 获取Pipeline Task, 因为Job Task是先保存在触发的回调, 这里获取的Pipeline Task是最新的
	descReq := task.NewDescribePipelineTaskRequest(in.Spec.PipelineTask)
	p, err := i.DescribePipelineTask(ctx, descReq)
	if err != nil {
		return nil, err
	}

	defer func() {
		// 更新当前任务的pipeline task状态
		i.mustUpdatePipelineStatus(ctx, p)

		// 如果JobTask正常执行, 则等待回调更新, 如果执行失败 则需要立即更新JobTask状态
		for index := range runErrorJobTasks {
			_, err = i.UpdateJobTaskStatus(ctx, runErrorJobTasks[index])
			if err != nil {
				p.MarkedFailed(err)
				i.mustUpdatePipelineStatus(ctx, p)
			}
		}
	}()

	// 更新Pipeline Task 运行时环境变量
	p.Status.RuntimeEnvs.Merge(in.Status.RuntimeEnvs.Params...)

	switch in.Status.Stage {
	case task.STAGE_PENDDING, task.STAGE_ACTIVE, task.STAGE_CANCELING:
		// Pipeline Task状态无变化
		return p, nil
	case task.STAGE_CANCELED:
		// 任务取消, pipeline 取消执行
		p.MarkedCanceled()
		return p, nil
	case task.STAGE_FAILED:
		// 任务执行失败, 更新Pipeline状态为失败
		if !in.Spec.IgnoreFailed {
			p.MarkedFailed(in.Status.MessageToError())
			return p, nil
		}
	case task.STAGE_SUCCEEDED:
		// 任务运行成功, pipeline继续执行
		i.log.Infof("task: %s run successed", in.Spec.TaskId)
	}

	// task执行成功或者忽略执行失败, 此时pipeline 仍然处于运行中, 需要获取下一个任务执行
	nexts, err := p.NextRun()
	if err != nil {
		p.MarkedFailed(err)
		return p, nil
	}

	// 如果没有需要执行的任务, Pipeline执行结束, 更新Pipeline状态为成功
	if nexts == nil || nexts.Len() == 0 {
		p.MarkedSuccess()
		return p, nil
	}

	// 如果有需要执行的JobTask, 继续执行
	for index := range nexts.Items {
		item := nexts.Items[index]
		// 如果任务执行成功则等待任务的回调更新任务状态
		// 如果任务执行失败, 直接更新任务状态
		t, err := i.RunJob(ctx, item.Spec)
		if err != nil {
			updateT := task.NewUpdateJobTaskStatusRequest(item.Spec.TaskId)
			updateT.UpdateToken = item.Spec.UpdateToken
			updateT.MarkError(err)
			runErrorJobTasks = append(runErrorJobTasks, updateT)
		} else {
			item.Status = t.Status
			item.Job = t.Job
		}
	}

	return p, nil
}

// 更新Pipeline状态
func (i *impl) updatePipelineStatus(ctx context.Context, in *task.PipelineTask) (*task.PipelineTask, error) {
	// 立即更新Pipeline Task状态
	i.updatePiplineTaskStatus(ctx, in)

	// 执行PipelineTask状态变更回调
	i.PipelineStatusChangedCallback(ctx, in)

	// 补充回调执行状态
	i.updatePiplineTaskStatus(ctx, in)
	return in, nil
}

func (i *impl) PipelineStatusChangedCallback(ctx context.Context, in *task.PipelineTask) {
	if !in.HasJobSpec() {
		return
	}

	if in.Status == nil {
		return
	}

	// WebHook回调
	webhooks := in.Pipeline.Spec.MatchedWebHooks(in.Status.Stage.String())
	i.hook.SendTaskStatus(ctx, webhooks, in)

	// 关注人通知回调
	for index := range in.Pipeline.Spec.MentionUsers {
		mu := in.Pipeline.Spec.MentionUsers[index]
		i.TaskMention(ctx, mu, in)
	}

	// 是否需要运行下一个Pipeline
	if in.Pipeline.Spec.NextPipeline != "" {
		// 如果有审核单则提交审核单, 没有则直接运行
		i.log.Debugf("next pipeline: %s", in.Pipeline.Spec.NextPipeline)
	}

	// 事件队列回调通知, 通知事件队列该事件触发的PipelineTask已经执行完成
	if in.IsComplete() {
		tReq := trigger.NewEventQueueTaskCompleteRequest(in.Meta.Id)
		bs, err := i.trigger.EventQueueTaskComplete(ctx, tReq)
		if err != nil {
			in.AddErrorEvent("触发队列回调失败: %s", err)
		} else {
			in.AddSuccessEvent("触发队列成功: %s", bs.BuildConfig.Spec.Name).SetDetail(bs.String())
		}
	}
}

// 更新Pipeline状态
func (i *impl) mustUpdatePipelineStatus(ctx context.Context, in *task.PipelineTask) {
	_, err := i.updatePipelineStatus(ctx, in)
	if err != nil {
		i.log.Error(err)
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
	// 运行中的流水线不运行删除, 先取消 才能删除
	if ins.IsRunning() {
		return nil, fmt.Errorf("流水线运行结束才能删除, 如果没结束, 请先取消再删除")
	}

	// 删除该Pipeline下所有的Job Task
	tasks := ins.Status.JobTasks()
	for index := range tasks.Items {
		t := tasks.Items[index]

		// 没有运行过的任务不需要清理
		if t.Status.Stage.Equal(task.STAGE_PENDDING) {
			continue
		}

		deleteReq := task.NewDeleteJobTaskRequest(t.Spec.TaskId)
		deleteReq.Force = in.Force
		_, err := i.DeleteJobTask(ctx, deleteReq)
		if err != nil {
			if !exception.IsNotFoundError(err) {
				return nil, err
			}
		}
	}

	if err := i.deletePipelineTask(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
