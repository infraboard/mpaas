package k8s

import (
	"context"
	"time"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/common/format"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func (r *K8sRunner) Run(ctx context.Context, in *task.RunTaskRequest) (
	*task.JobTaskStatus, error) {
	runnerParams := in.Params.K8SJobRunnerParams()
	cReq := cluster.NewDescribeClusterRequest(runnerParams.ClusterId)
	c, err := r.cluster.DescribeCluster(ctx, cReq)
	if err != nil {
		return nil, err
	}
	k8sClient, err := c.Client()
	if err != nil {
		return nil, err
	}
	r.k8sClient = k8sClient
	r.log.Infof("get k8s cluster ok, %s [%s]", c.Spec.Name, c.Meta.Id)

	obj := new(v1.Job)
	jobYamlSpec := in.RenderJobSpec()
	r.log.Debugf("job rendered yaml spec: %s", jobYamlSpec)
	if err := yaml.Unmarshal([]byte(jobYamlSpec), obj); err != nil {
		return nil, err
	}

	// 处理系统变量
	if err := r.HanleSystemVariable(ctx, in.Params, obj); err != nil {
		return nil, err
	}

	// 修改任务名称
	obj.Name = in.Name
	obj.Namespace = runnerParams.Namespace

	// Job注入标签
	workload.InjectJobLabels(obj, in.Labels)
	// Job注入注解
	workload.InjectJobAnnotations(obj, in.Annotations())
	// 给Job容器注入环境变量
	workload.InjectPodEnvVars(&obj.Spec.Template.Spec, in.Params.EnvVars())

	status := task.NewJobTaskStatus()
	status.MarkedRunning()

	// 运行时环境变量注入
	err = r.PrepareRuntime(ctx, k8sClient, in, obj, status)
	if err != nil {
		return nil, err
	}
	defer r.CleanUpRuntime(ctx, k8sClient, in, status)

	// 执行Job
	if !in.DryRun {
		r.log.Debug("run job yaml: %s", format.MustToYaml(obj))
		obj, err = k8sClient.WorkLoad().CreateJob(ctx, obj)
		if err != nil {
			return nil, err
		}
	}

	objYaml, err := yaml.Marshal(obj)
	if err != nil {
		return nil, err
	}
	status.Detail = string(objYaml)
	return status, nil
}

// 准备
func (r *K8sRunner) PrepareRuntime(
	ctx context.Context,
	k8sClient *k8s.Client,
	in *task.RunTaskRequest,
	obj *v1.Job,
	status *task.JobTaskStatus,
) error {

	// 创建一个configmap 用于收集Task运行时的中间信息(以环境变量的方式)
	runtimeEnvConfigMap := in.RuntimeEnvConfigMap(task.CONFIG_MAP_RUNTIME_ENV_MOUNT_PATH)
	runtimeEnvConfigMap.Namespace = in.Params.K8SJobRunnerParams().Namespace
	r.log.Infof("create job runtime env configmap: %s", runtimeEnvConfigMap.Name)
	err := k8sClient.Config().FindOrCreateConfigMap(ctx, runtimeEnvConfigMap)
	if err != nil {
		return err
	}

	// 把configmap 注入为卷进行挂载,
	workload.InjectPodConfigMapVolume(&obj.Spec.Template.Spec, runtimeEnvConfigMap)

	// 更新临时资源等待释放
	tr := task.NewTemporaryResource(runtimeEnvConfigMap.Kind, runtimeEnvConfigMap.Name)
	status.AddTemporaryResource(tr)
	return nil
}

// 临时资源回收与Runtime Env更新
// 不用清理, PipelineTask结束时统一清理
func (r *K8sRunner) CleanUpRuntime(
	ctx context.Context,
	k8sClient *k8s.Client,
	in *task.RunTaskRequest,
	status *task.JobTaskStatus) {
	// 运行结束 从config map中读取Env, 并更新到Task状态中去
	cmName := task.NewJobTaskEnvConfigMapName(in.Params.GetJobTaskId())
	ns := in.Params.K8SJobRunnerParams().Namespace
	req := meta.NewGetRequest(cmName).WithNamespace(ns)
	runtimeEnvConfigMap, err := k8sClient.Config().GetConfigMap(ctx, req)
	if err != nil {
		r.log.Errorf("get config map error, %s", err)
		return
	}

	// 解析并更新Runtime Env
	data := runtimeEnvConfigMap.BinaryData[task.CONFIG_MAP_RUNTIME_ENV_KEY]
	envs, err := task.ParseRuntimeEnvFromBytes(data)
	if err != nil {
		r.log.Errorf("parse env data error, %s", err)
		return
	}
	status.RuntimeEnvs = envs

	// 清除临时挂载的configmap
	err = k8sClient.Config().DeleteConfigMap(ctx, meta.NewDeleteRequest(cmName).WithNamespace(ns))
	if err != nil {
		r.log.Errorf("delete config map error, %s", err)
		return
	}

	r.log.Info("delete job runtime env configmap: %s", cmName)
	tr := status.GetTemporaryResource("configmap", cmName)
	if tr != nil {
		tr.ReleaseAt = time.Now().Unix()
	}
}
