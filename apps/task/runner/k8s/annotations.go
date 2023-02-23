package k8s

import (
	"context"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

// 更加注解注入信息

func (r *K8sRunner) HanleSystemVariable(ctx context.Context, in *job.VersionedRunParam, job *v1.Job) {
	r.handleDeployConfig(ctx, in.GetDeployConfigId(), job)
}

// 查询部署配置, 注入相关变量
func (r *K8sRunner) handleDeployConfig(ctx context.Context, deployConfigId string, job *v1.Job) error {
	if deployConfigId == "" {
		return nil
	}

	// 查询部署配置
	dc, err := r.deploy.DescribeDeployConfig(ctx, deploy.NewDescribeDeployConfigRequest(deployConfigId))
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
		secret := c.KubeConfSecret()
		secret.Annotations[workload.SECRET_MOUNT_ANNOTATION_KEY] = "/.kube"
		checkReq := meta.NewGetRequest(secret.Name).WithNamespace(job.Namespace)
		_, err = r.k8sClient.Config().GetSecret(ctx, checkReq)
		if errors.IsNotFound(err) {
			_, err := r.k8sClient.Config().CreateSecret(ctx, secret)
			if err != nil {
				return err
			}
		}
		workload.InjectPodSecretVolume(&job.Spec.Template.Spec, secret)
	case deploy.TYPE_HOST:
		// 主机部署需要注入的信息
	}

	return nil
}
