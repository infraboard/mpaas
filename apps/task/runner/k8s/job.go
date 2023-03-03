package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/common/format"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/config"
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

	// 执行Job
	if !in.DryRun {
		r.log.Debugf("run job yaml: %s", format.MustToYaml(obj))
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

// Task运行时 需要提前准备一些资源
// 资源的清理在Task 状态更新时执行, 不在这里执行,
// 这里任务的异步执行的, k8s job是异步执行
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
	tr := task.NewTemporaryResource(
		config.CONFIG_KIND_CONFIG_MAP.String(),
		runtimeEnvConfigMap.Name,
	)
	tr.Detail = format.MustToYaml(runtimeEnvConfigMap)
	status.AddTemporaryResource(tr)
	return nil
}
