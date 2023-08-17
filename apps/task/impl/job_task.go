package impl

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/runner"
	k8s_provider "github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/config"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"
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

	// 忽略执行
	if in.Enabled() {
		// 查询需要执行的Job
		req := job.NewDescribeJobRequestByName(in.JobName)

		j, err := i.job.DescribeJob(ctx, req)
		if err != nil {
			return nil, err
		}
		ins.Job = j
		i.log.Infof("describe job success, %s[%s]", j.Spec.Name, j.Meta.Id)

		// 脱敏参数动态还原
		in.RunParams.RestoreSensitive(j.Spec.RunParam)

		// 合并允许参数(Job里面有默认值), 并检查参数合法性
		// 注意Param的合并是有顺序的，也就是参数优先级(低-->高):
		// 1. 系统变量(默认禁止修改)
		// 2. job默认变量
		// 3. job运行变量
		// 4. pipeline 运行变量
		// 5. pipeline 运行时变量
		params := job.NewRunParamSet()
		params.Add(ins.SystemRunParam()...)
		params.Add(j.Spec.RunParam.Params...)
		params.Merge(in.RunParams.Params...)
		err = i.LoadPipelineRunParam(ctx, params)
		if err != nil {
			return nil, err
		}

		// 校验参数合法性
		err = params.Validate()
		if err != nil {
			return nil, fmt.Errorf("校验任务【%s】参数错误, %s", j.Spec.Name, err)
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
		status.RunParams = params
		ins.Status = status

		// 添加搜索标签
		ins.BuildSearchLabel()
	}

	// 保存任务
	updateOpt := options.Update()
	updateOpt.SetUpsert(true)
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": ins}, updateOpt); err != nil {
		return nil, exception.NewInternalServerError("inserted a job task document error, %s", err)
	}
	return ins, nil
}

func (r *impl) GetK8sClient(ctx context.Context, req *job.K8SJobRunnerParams) (*k8s_provider.Client, error) {
	if req.KubeConfig != "" {
		return req.Client()
	}

	cReq := k8s.NewDescribeClusterRequest(req.ClusterId)
	c, err := r.cluster.DescribeCluster(ctx, cReq)
	if err != nil {
		return nil, err
	}
	r.log.Infof("get k8s cluster ok, %s [%s]", c.Spec.Name, c.Meta.Id)
	return c.Client()
}

// 加载Pipeline 提供的运行时参数
func (i *impl) LoadPipelineRunParam(ctx context.Context, params *job.RunParamSet) error {
	pipelineTaskId := params.GetPipelineTaskId()
	if pipelineTaskId == "" {
		return nil
	}
	// 查询出Pipeline
	pt, err := i.DescribePipelineTask(ctx, task.NewDescribePipelineTaskRequest(pipelineTaskId))
	if err != nil {
		return err
	}

	// 合并PipelineTask传入的变量参数
	params.Merge(pt.Params.RunParams...)
	// 合并PipelineTask的运行时参数, Task运行时更新的
	params.Merge(pt.RuntimeRunParams()...)
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

func (i *impl) CheckAllowUpdate(ctx context.Context, ins *task.JobTask, token string, force bool) error {
	// 校验更新合法性
	err := ins.ValidateToken(token)
	if err != nil {
		return err
	}

	// 修改任务状态
	if !force && ins.Status.IsComplete() {
		return exception.NewBadRequest("状态[%s], 已经结束的任务不能更新状态", ins.Status.Stage)
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
	err = i.CheckAllowUpdate(ctx, ins, in.UpdateToken, in.Force)
	if err != nil {
		return nil, err
	}
	ins.Status.UpdateOutput(in)

	// 只更新任务状态
	if _, err := i.jcol.UpdateByID(ctx, ins.Spec.TaskId, bson.M{"$set": bson.M{"status": ins.Status}}); err != nil {
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
	err = i.CheckAllowUpdate(ctx, ins, in.UpdateToken, in.ForceUpdateStatus)
	if err != nil {
		return nil, err
	}

	// 状态更新
	preStatus := ins.Status.Stage
	ins.Status.UpdateStatus(in)

	// Job Task状态变更回调
	i.JobTaskStatusChangedCallback(ctx, ins)

	// 更新数据库
	if err := i.updateJobTask(ctx, ins); err != nil {
		return nil, err
	}

	// Pipeline Task 状态变更回调
	if ins.Spec.PipelineTask != "" {
		// 如果状态未变化, 不触发流水线更新
		if !in.ForceTriggerPipeline && preStatus.Equal(in.Stage) {
			i.log.Debugf("task %s status not changed: %s, skip update pipeline", in.Id, in.Stage)
			return ins, nil
		}
		_, err := i.PipelineTaskStatusChanged(ctx, ins)
		if err != nil {
			return nil, err
		}
	}
	return ins, nil
}

func (i *impl) JobTaskStatusChangedCallback(ctx context.Context, in *task.JobTask) {
	if !in.HasJobSpec() {
		return
	}

	if in.Status == nil {
		return
	}

	// WebHook回调
	webhooks := in.Spec.MatchedWebHooks(in.Status.Stage.String())
	i.hook.SendTaskStatus(ctx, webhooks, in)

	// 关注人通知回调
	for index := range in.Spec.MentionUsers {
		mu := in.Spec.MentionUsers[index]
		i.TaskMention(ctx, mu, in)
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
		k8sParams := in.Status.RunParams.K8SJobRunnerParams()
		k8sClient, err := i.GetK8sClient(
			ctx,
			k8sParams,
		)
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

// 查询Task日志
func (i *impl) WatchJobTaskLog(in *task.WatchJobTaskLogRequest, stream task.JobRPC_WatchJobTaskLogServer) error {
	writer := NewWatchJobTaskLogServerWriter(stream)
	writer.WriteMessagef("正在查询Task[%s]的日志 请稍等...", in.TaskId)

	// 等待Task的Pod正常启动
	t, err := i.WaitPodLogReady(stream.Context(), in, writer)
	if err != nil {
		return err
	}

	switch t.Job.Spec.RunnerType {
	case job.RUNNER_TYPE_K8S_JOB:
		k8sParams := t.Job.Spec.RunParam.K8SJobRunnerParams()
		k8sClient, err := i.GetK8sClient(stream.Context(), k8sParams)
		if err != nil {
			return err
		}

		// 找到Job执行的Pod
		podReq := meta.NewListRequest().
			SetNamespace(k8sParams.Namespace).
			SetLabelSelector(meta.NewLabelSelector().Add("job-name", t.Spec.TaskId))
		pods, err := k8sClient.WorkLoad().ListPod(stream.Context(), podReq)
		if err != nil {
			return err
		}
		if len(pods.Items) == 0 {
			return fmt.Errorf("job's pod not found by lable job-name=%s", t.Spec.TaskId)
		}

		req := workload.NewWatchConainterLogRequest()
		req.PodName = pods.Items[0].Name
		req.Namespace = k8sParams.Namespace
		req.Container = in.ContainerName
		r, err := k8sClient.WorkLoad().WatchConainterLog(stream.Context(), req)
		if err != nil {
			return err
		}
		defer r.Close()

		// copy日志流
		_, err = io.Copy(writer, r)
		return err
	}

	return nil
}

func (i *impl) WaitPodLogReady(
	ctx context.Context,
	in *task.WatchJobTaskLogRequest,
	writer *WatchJobTaskLogServerWriter,
) (*task.JobTask, error) {
	maxRetryCount := 0
WAIT_TASK_ACTIVE:
	// 查询Task信息
	t, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.TaskId))
	if err != nil {
		return nil, err
	}

	pod, err := t.Status.GetLatestPod()
	if err != nil {
		return nil, err
	}

	if maxRetryCount < 30 {
		if pod == nil {
			writer.WriteMessagef("任务当前状态: [%s], Pod创建中...", t.Status.Stage)
		} else {
			writer.WriteMessagef("任务当前状态: [%s], Pod状态: [%s], 等待任务启动中...", t.Status.Stage, pod.Status.Phase)
			// Job状态运行成功，返回
			if pod.Status.Phase != "Pending" {
				return t, nil
			}
		}

		time.Sleep(2 * time.Second)
		maxRetryCount++
		goto WAIT_TASK_ACTIVE
	}

	return t, nil
}

// Task Debug
func (i *impl) JobTaskDebug(ctx context.Context, in *task.JobTaskDebugRequest) {
	term := in.WebTerminal()

	// 查询Task信息
	t, err := i.DescribeJobTask(ctx, task.NewDescribeJobTaskRequest(in.TaskId))
	if err != nil {
		term.Failed(err)
		return
	}

	switch t.Job.Spec.RunnerType {
	case job.RUNNER_TYPE_K8S_JOB:
		k8sParams := t.Status.RunParams.K8SJobRunnerParams()
		k8sClient, err := i.GetK8sClient(ctx, k8sParams)
		if err != nil {
			term.Failed(fmt.Errorf("初始化k8s客户端失败, %s", err))
			return
		}
		if err != nil {
			term.Failed(err)
			return
		}

		// 找到Job执行的Pod
		term.WriteTextln("正在查询Job Task【%s】运行的Pod", t.Spec.TaskId)
		podReq := meta.NewListRequest().
			SetNamespace(k8sParams.Namespace).
			SetLabelSelector(meta.NewLabelSelector().Add("job-name", t.Spec.TaskId))
		pods, err := k8sClient.WorkLoad().ListPod(ctx, podReq)
		if err != nil {
			term.Failed(err)
			return
		}
		if len(pods.Items) == 0 {
			term.Failed(fmt.Errorf("job's pod not found by lable job-name=%s", t.Spec.TaskId))
			return
		}

		targetCopyPod := pods.Items[0]
		term.WriteTextln("Job Task【%s】位于Namespace: %s, PodName: %s",
			t.Spec.TaskId, targetCopyPod.Namespace,
			targetCopyPod.Name,
		)

		req := in.CopyPodRunRequest(k8sParams.Namespace, targetCopyPod.Name)
		req.SetAttachTerminal(term)
		req.Remove = true

		_, err = k8sClient.WorkLoad().CopyPodRun(ctx, req)
		if err != nil {
			term.Failed(err)
			return
		}
	default:
		term.Failed(fmt.Errorf("unknonw runner type %s", t.Job.Spec.RunnerType))
		return
	}
}

func NewWatchJobTaskLogServerWriter(
	stream task.JobRPC_WatchJobTaskLogServer) *WatchJobTaskLogServerWriter {
	return &WatchJobTaskLogServerWriter{
		stream: stream,
		buf:    task.NewJobTaskStreamReponse(),
	}
}

type WatchJobTaskLogServerWriter struct {
	stream task.JobRPC_WatchJobTaskLogServer
	buf    *task.JobTaskStreamReponse
}

func (w *WatchJobTaskLogServerWriter) Write(p []byte) (n int, err error) {
	w.buf.Data = p
	err = w.stream.Send(w.buf)
	if err != nil {
		return 0, err
	}
	w.buf.ReSet()
	return len(p), nil
}

func (w *WatchJobTaskLogServerWriter) WriteMessagef(format string, a ...any) {
	_, err := w.Write([]byte(fmt.Sprintf(format+"\n", a...)))
	if err != nil {
		zap.L().Errorf("write message error, %s", err)
	}
}
