package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func (r *K8sRunner) Run(ctx context.Context, in *task.RunTaskRequest) (*task.JobTaskStatus, error) {
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

	obj := new(v1.Job)
	if err := yaml.Unmarshal([]byte(in.JobSpec), obj); err != nil {
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
	// 给容器注入环境变量
	workload.InjectPodEnvVars(&obj.Spec.Template.Spec, in.Params.EnvVars())

	status := task.NewJobTaskStatus()
	status.MarkedRunning()

	// 执行Job
	if !in.DryRun {
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
