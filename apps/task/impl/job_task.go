package impl

import (
	"context"
	"fmt"
	"time"

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
	"github.com/infraboard/mpaas/provider/k8s/config"
	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func (i *impl) RunJob(ctx context.Context, in *pipeline.RunJobRequest) (
	*task.JobTask, error) {
	if in.TaskId != "" {
		// 如果任务重新运行, 需要等待之前的任务结束后才能执行
		isActive, err := i.CheckJotTaskIsActive(ctx, in.TaskId)
		if err != nil {
			return nil, err
		}
		if isActive {
			return nil, exception.NewConflict("任务: %s 当前处于运行中, 需要等待运行结束后才能执行", in.TaskId)
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
	i.log.Infof("describe job success, %s[%s]", j.Spec.Name, j.Meta.Id)

	// 合并允许参数(Job里面有默认值), 并检查参数合法性
	params := j.GetVersionedRunParam(in.GetRunParamsVersion())
	if params == nil {
		return nil, fmt.Errorf("job %s version: %s not found, allow version: %s",
			j.Spec.Name,
			in.GetRunParamsVersion(),
			j.AllowVersions(),
		)
	}
	params.Merge(in.RunParams)
	params.Add(ins.SystemVariable()...)
	err = i.LoadRuntimeEnvs(ctx, in.RunParams.GetPipelineTaskId(), params)
	if err != nil {
		return nil, err
	}
	err = params.Validate()
	if err != nil {
		return nil, err
	}
	i.log.Infof("params check ok, %s", params)

	// 获取执行器执行
	r := runner.GetRunner(j.Spec.RunnerType)
	runReq := task.NewRunTaskRequest(ins.Spec.TaskId, j.Spec.RunnerSpec, params)
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
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": ins}, updateOpt); err != nil {
		return nil, exception.NewInternalServerError("inserted a job task document error, %s", err)
	}
	return ins, nil
}

// 加载Pipeline 提供的Runtime env
func (i *impl) LoadRuntimeEnvs(ctx context.Context, pipelineId string, params *job.VersionedRunParam) error {
	if pipelineId == "" {
		return nil
	}
	pt, err := i.DescribePipelineTask(ctx, task.NewDescribePipelineTaskRequest(pipelineId))
	if err != nil {
		return err
	}
	params.UpdateFromEnvs(pt.RuntimeEnvVars())
	return nil
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

func (i *impl) CheckAllowUpdate(ctx context.Context, ins *task.JobTask, token string) error {
	// 校验更新合法性
	err := ins.ValidateToken(token)
	if err != nil {
		return err
	}

	// 修改任务状态
	if ins.Status.IsComplete() {
		return exception.NewBadRequest("已经结束的任务不能更新状态")
	}

	return nil
}

// 更新任务运行结果
func (i *impl) UpdateJobTaskOutput(ctx context.Context, in *task.UpdateJobTaskOutputRequest) (
	*task.JobTask, error) {
	ins, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}

	// 校验更新合法性
	err = i.CheckAllowUpdate(ctx, ins, in.UpdateToken)
	if err != nil {
		return nil, err
	}
	ins.Status.UpdateOutput(in)

	// 更新数据库
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": ins}); err != nil {
		return nil, exception.NewInternalServerError("update task(%s) document error, %s",
			in.Id, err)
	}

	return ins, nil
}

// 更新Job状态
func (i *impl) UpdateJobTaskStatus(ctx context.Context, in *task.UpdateJobTaskStatusRequest) (
	*task.JobTask, error) {
	ins, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}

	// 校验更新合法性
	err = i.CheckAllowUpdate(ctx, ins, in.UpdateToken)
	if err != nil {
		return nil, err
	}
	ins.Status.UpdateStatus(in)

	// 任务状态变化处理
	// i.StatusChangedHook(ctx, ins)

	// 更新数据库
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": ins}); err != nil {
		return nil, exception.NewInternalServerError("update task(%s) document error, %s",
			in.Id, err)
	}

	// Pipeline Task 状态变更回调
	if ins.Spec.PipelineTask != "" {
		_, err := i.PipelineTaskStatusChanged(ctx, ins)
		if err != nil {
			return nil, err
		}
	}
	return ins, nil
}

func (i *impl) StatusChangedHook(ctx context.Context, in *task.JobTask) {
	if !in.HasJobSpec() {
		return
	}

	switch in.Job.Spec.RunnerType {
	case job.RUNNER_TYPE_K8S_JOB:
		jobParams := in.Job.GetVersionedRunParam(in.Spec.RunParams.Version)
		if jobParams == nil {
			in.Status.AddErrorEvent("job version params not found")
			return
		}
		k8sParams := jobParams.K8SJobRunnerParams()

		descReq := cluster.NewDescribeClusterRequest(k8sParams.ClusterId)
		c, err := i.cluster.DescribeCluster(ctx, descReq)
		if err != nil {
			in.Status.AddErrorEvent("find k8s cluster error, %s", err)
			return
		}

		k8sClient, err := c.Client()
		if err != nil {
			in.Status.AddErrorEvent("init k8s client error, %s", err)
			return
		}

		// 读取挂载的runtime configmap
		cmName := task.NewJobTaskEnvConfigMapName(in.Spec.TaskId)
		req := meta.NewGetRequest(cmName).WithNamespace(k8sParams.Namespace)
		runtimeEnvConfigMap, err := k8sClient.Config().GetConfigMap(ctx, req)
		if err != nil {
			in.Status.AddErrorEvent("get config map error, %s", err)
			return
		}
		// 解析并更新Runtime Env
		data := runtimeEnvConfigMap.BinaryData[task.CONFIG_MAP_RUNTIME_ENV_KEY]
		envs, err := task.ParseRuntimeEnvFromBytes(data)
		if err != nil {
			in.Status.AddErrorEvent("parse env data error, %s", err)
			return
		}
		in.Status.RuntimeEnvs = envs
	}
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

	// 清理Job关联的临时资源
	err = i.CleanTaskResource(ctx, ins)
	if err != nil {
		if !in.Force {
			return nil, err
		}
		i.log.Warnf("force delete, but has error, %s", err)
	}

	// 删除本地记录
	_, err = i.jcol.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) CleanTaskResource(ctx context.Context, in *task.JobTask) error {
	if !in.HasJobSpec() {
		return fmt.Errorf("job spec is nil")
	}

	switch in.Job.Spec.RunnerType {
	case job.RUNNER_TYPE_K8S_JOB:
		jobParams, err := in.GetVersionedRunParam()
		if err != nil {
			return err
		}
		k8sParams := jobParams.K8SJobRunnerParams()

		descReq := cluster.NewDescribeClusterRequest(k8sParams.ClusterId)
		c, err := i.cluster.DescribeCluster(ctx, descReq)
		if err != nil {
			return fmt.Errorf("find k8s cluster error, %s", err)
		}

		k8sClient, err := c.Client()
		if err != nil {
			return err
		}

		// 清除临时挂载的configmap
		for i := range in.Status.TemporaryResources {
			resource := in.Status.TemporaryResources[i]
			if resource.IsReleased() {
				continue
			}
			switch resource.Kind {
			case config.CONFIG_KIND_CONFIG_MAP.String():
				cmDeleteReq := meta.NewDeleteRequest(resource.Name).WithNamespace(k8sParams.Namespace)
				err = k8sClient.Config().DeleteConfigMap(ctx, cmDeleteReq)
				if err != nil {
					return fmt.Errorf("delete config map error, %s", err)
				}
				in.Status.AddEvent(task.EVENT_LEVEL_DEBUG, "delete job runtime env configmap: %s", resource.Name)
				resource.ReleaseAt = time.Now().Unix()
			}
		}

		// 清理Job
		detail := in.GetStatusDetail()
		if detail == "" {
			return fmt.Errorf("no k8s job found in status detail")
		}

		obj := new(v1.Job)
		if err := yaml.Unmarshal([]byte(detail), obj); err != nil {
			return err
		}

		req := meta.NewDeleteRequest(obj.Name)
		req.Namespace = obj.Namespace
		err = k8sClient.WorkLoad().DeleteJob(ctx, req)
		if err != nil {
			return fmt.Errorf("delete k8s job error, %s", err)
		}
		return err
	}

	return nil
}
