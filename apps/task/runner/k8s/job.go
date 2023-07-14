package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/job"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/common/format"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func (r *K8sRunner) Run(ctx context.Context, in *task.RunTaskRequest) (
	*task.JobTaskStatus, error) {

	// 获取k8s集群参数
	runnerParams := in.Params.K8SJobRunnerParams()
	k8sClient, err := r.GetK8sClient(ctx, runnerParams)
	if err != nil {
		return nil, err
	}

	// 获取Job定义
	obj := new(v1.Job)
	jobYamlSpec := in.RenderJobSpec()
	r.log.Debugf("job rendered yaml spec: %s", jobYamlSpec)
	if err := yaml.Unmarshal([]byte(jobYamlSpec), obj); err != nil {
		return nil, err
	}

	// 处理系统变量
	hr := NewHanleSystemVariableRequest(k8sClient, in.Params, obj)
	if err := r.HanleSystemVariable(ctx, hr); err != nil {
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

func (r *K8sRunner) GetK8sClient(ctx context.Context, req *job.K8SJobRunnerParams) (*k8s.Client, error) {
	if req.KubeConfig != "" {
		return req.Client()
	}

	cReq := cluster.NewDescribeClusterRequest(req.ClusterId)
	c, err := r.cluster.DescribeCluster(ctx, cReq)
	if err != nil {
		return nil, err
	}
	r.log.Infof("get k8s cluster ok, %s [%s]", c.Spec.Name, c.Meta.Id)
	return c.Client()
}
