package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/runner"
	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func (i *impl) RunJob(ctx context.Context, in *pipeline.RunJobRequest) (
	*task.JobTask, error) {
	if in.Id != "" {
		// 如果任务重新运行, 需要等待之前的任务结束后才能执行
		isActive, err := i.CheckJotTaskIsActive(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		if isActive {
			return nil, exception.NewConflict("任务: %s 当前处于运行中, 需要等待运行结束后才能执行", in.Id)
		}
	}

	ins := task.NewJobTask(in)

	// 查询需要执行的Job
	req := job.NewDescribeJobRequest(in.JobName)
	j, err := i.job.DescribeJob(ctx, req)
	if err != nil {
		return nil, err
	}
	ins.Job = j

	// 合并允许参数(Job里面有默认值), 并检查参数合法性
	params := j.GetVersionedRunParam(in.Params.Version)
	params.Merge(in.Params)
	params.Add(ins.SystemVariable()...)
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// 获取执行器执行
	r := runner.GetRunner(j.Spec.RunnerType)
	runReq := task.NewRunTaskRequest(ins.Spec.Id, j.Spec.RunnerSpec, params)
	runReq.DryRun = in.DryRun
	runReq.Labels = in.Labels
	runReq.ManualUpdateStatus = j.Spec.ManualUpdateStatus
	status, err := r.Run(ctx, runReq)
	if err != nil {
		return nil, err
	}
	ins.Status = status

	// 3. 保存任务
	updateOpt := options.Update()
	updateOpt.SetUpsert(true)
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.Id, bson.M{"$set": ins}, updateOpt); err != nil {
		return nil, exception.NewInternalServerError("inserted a job task document error, %s", err)
	}
	return ins, nil
}

// 判断任务是否还处于运行中
func (i *impl) CheckJotTaskIsActive(ctx context.Context, jotTaskId string) (bool, error) {
	ins, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(jotTaskId))
	if err != nil && !exception.IsNotFoundError(err) {
		return false, err
	}

	return ins.Status.Stage.Equal(task.STAGE_ACTIVE), nil

}

func (i *impl) JobTaskBatchSave(ctx context.Context, in *task.JobTaskSet) error {
	if _, err := i.jcol.InsertMany(ctx, in.ToDocs()); err != nil {
		return exception.NewInternalServerError("inserted job tasks document error, %s", err)
	}
	return nil
}

func (i *impl) QueryJobTask(ctx context.Context, in *task.QueryJobTaskRequest) (
	*task.JobTaskSet, error) {
	r := newQueryRequest(in)
	resp, err := i.jcol.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := task.NewJobTaskSet()
	// 循环
	for resp.Next(ctx) {
		ins := task.NewDefaultJobTask()
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

// 更新Job状态
func (i *impl) UpdateJobTaskStatus(ctx context.Context, in *task.UpdateJobTaskStatusRequest) (
	*task.JobTask, error) {
	ins, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}

	// 修改任务状态
	if ins.Status.IsComplete() {
		return nil, exception.NewBadRequest("已经结束的任务不能更新状态")
	}
	ins.Status.Update(in)

	// 更新数据库
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.Id, bson.M{"$set": ins}); err != nil {
		return nil, exception.NewInternalServerError("update task(%s) document error, %s",
			in.Id, err)
	}

	// Pipeline Task 状态变更回调
	if ins.Spec.PipelineTask != "" {
		i.PipelineTaskStatusChanged(ctx, ins)
	}
	return ins, nil
}

// 任务执行详情
func (i *impl) DescribeJobTask(ctx context.Context, in *task.DescribeJobTaskRequest) (
	*task.JobTask, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := task.NewDefaultJobTask()
	if err := i.jcol.FindOne(ctx, bson.M{"_id": in.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("job task %s not found", in.Id)
		}

		return nil, exception.NewInternalServerError("find job task %s error, %s", in.Id, err)
	}
	return ins, nil
}

// 删除任务
func (i *impl) DeleteJobTask(ctx context.Context, in *task.DeleteJobTaskRequest) (
	*task.JobTask, error) {
	ins, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}

	// 任务清理
	if ins.HasJobSpec() {
		switch ins.Job.Spec.RunnerType {
		case job.RUNNER_TYPE_K8S_JOB:
			err = i.deleteK8sJob(ctx, ins)
			if err != nil {
				return nil, fmt.Errorf("delete k8s job error, %s", err)
			}
		}
	}

	// 删除本地记录
	_, err = i.jcol.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// 删除k8s中对应的job
func (i *impl) deleteK8sJob(ctx context.Context, ins *task.JobTask) error {
	jobParams := ins.Job.GetVersionedRunParam(ins.Spec.Params.Version)
	if jobParams == nil {
		return fmt.Errorf("job version params not found")
	}

	k8sParams := jobParams.K8SJobRunnerParams()
	c, err := i.cluster.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(k8sParams.ClusterId))
	if err != nil {
		return err
	}
	k8sClient, err := c.Client()
	if err != nil {
		return err
	}

	detail := ins.GetStatusDetail()
	if detail == "" {
		return fmt.Errorf("no k8s job found in status detail")
	}

	fmt.Println(detail)
	obj := new(v1.Job)
	if err := yaml.Unmarshal([]byte(detail), obj); err != nil {
		return err
	}

	fmt.Println(obj)

	req := meta.NewDeleteRequest(obj.Name)
	req.Namespace = obj.Namespace
	return k8sClient.WorkLoad().DeleteJob(ctx, req)
}
