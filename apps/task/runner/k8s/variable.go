package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
)

// 对系统变量进行处理

func (r *K8sRunner) HanleSystemVariable(ctx context.Context, in *job.VersionedRunParam, job *v1.Job) error {
	return r.handleDeployment(ctx, in, job)
}

// 查询部署配置, 注入相关变量
func (r *K8sRunner) handleDeployment(ctx context.Context, in *job.VersionedRunParam, job *v1.Job) error {
	DeploymentId := in.GetDeploymentId()

	if DeploymentId == "" {
		return nil
	}

	// 查询部署配置
	dc, err := r.deploy.DescribeDeployment(ctx, deploy.NewDescribeDeploymentRequest(DeploymentId))
	if err != nil {
		return err
	}

	switch dc.Spec.Type {
	case deploy.TYPE_KUBERNETES:
		// 容器部署需要注入的信息
		descReq := cluster.NewDescribeClusterRequest(dc.Spec.K8STypeConfig.ClusterId)
		c, err := r.cluster.DescribeCluster(ctx, descReq)
		if err != nil {
			return err
		}
		// 如果没有则创建Secret 并挂载到/.kube, 注意secret要声明挂载注解
		secret := c.KubeConfSecret("/.kube")
		secret.Namespace = in.K8SJobRunnerParams().Namespace
		err = r.k8sClient.Config().FindOrCreate(ctx, secret)
		if err != nil {
			return err
		}
		workload.InjectPodSecretVolume(&job.Spec.Template.Spec, secret)

		// 注入系统变量
		variables, err := dc.SystemVariable()
		if err != nil {
			return err
		}
		// 给容器注入环境变量
		workload.InjectPodEnvVars(&job.Spec.Template.Spec, variables)
	case deploy.TYPE_HOST:
		// 主机部署需要注入的信息
	}

	return nil
}
