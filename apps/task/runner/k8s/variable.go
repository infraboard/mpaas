package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/job"
	k8scluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
	core_v1 "k8s.io/api/core/v1"
)

func NewHanleSystemVariableRequest(
	client *k8s.Client,
	params *job.RunParamSet,
	job *v1.Job,
) *HanleSystemVariableRequest {
	return &HanleSystemVariableRequest{
		client: client,
		params: params,
		job:    job,
	}
}

// 对系统变量进行处理
type HanleSystemVariableRequest struct {
	client *k8s.Client
	params *job.RunParamSet
	job    *v1.Job
}

func (r *HanleSystemVariableRequest) PodSpec() *core_v1.PodSpec {
	return &r.job.Spec.Template.Spec
}

func (r *K8sRunner) HanleSystemVariable(ctx context.Context, in *HanleSystemVariableRequest) error {
	return r.handleDeployment(ctx, in)
}

// 查询部署配置, 注入相关变量
func (r *K8sRunner) handleDeployment(ctx context.Context, in *HanleSystemVariableRequest) error {
	DeploymentId := in.params.GetDeploymentId()

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

		params := in.params.K8SJobRunnerParams()
		var secret *core_v1.Secret
		if params.KubeConfig != "" {
			secret = params.KubeConfSecret(in.params.GetJobId(), "/.kube")
		} else {
			// 容器部署需要注入的信息
			descReq := k8scluster.NewDescribeClusterRequest(dc.Spec.K8STypeConfig.ClusterId)
			c, err := r.cluster.DescribeCluster(ctx, descReq)
			if err != nil {
				return err
			}
			// 如果没有则创建Secret 并挂载到/.kube, 注意secret要声明挂载注解
			secret = c.KubeConfSecret("/.kube")
		}

		secret.Namespace = params.Namespace
		err = in.client.Config().FindOrCreateSecret(ctx, secret)
		if err != nil {
			return err
		}
		workload.InjectPodSecretVolume(in.PodSpec(), secret)

		// 注入系统变量
		variables, err := dc.SystemVariable()
		if err != nil {
			return err
		}
		// 给容器注入环境变量
		workload.InjectPodEnvVars(in.PodSpec(), variables)
	case deploy.TYPE_HOST:
		// 主机部署需要注入的信息
	}

	return nil
}
