package workload

import (
	"context"
	"time"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
)

func (c *Workload) ListDeployment(ctx context.Context, req *meta.ListRequest) (*appsv1.DeploymentList, error) {
	ds, err := c.appsv1.Deployments(req.Namespace).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}
	if req.SkipManagedFields {
		for i := range ds.Items {
			ds.Items[i].ManagedFields = nil
		}
	}
	return ds, nil
}

func (c *Workload) GetDeployment(ctx context.Context, req *meta.GetRequest) (*appsv1.Deployment, error) {
	d, err := c.appsv1.Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	d.APIVersion = "apps/v1"
	d.Kind = "Deployment"
	return d, nil
}

func (c *Workload) WatchDeployment(ctx context.Context, req *appsv1.Deployment) (watch.Interface, error) {
	return c.appsv1.Deployments(req.Namespace).Watch(ctx, metav1.ListOptions{})
}

func (c *Workload) CreateDeployment(ctx context.Context, req *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.appsv1.Deployments(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Workload) UpdateDeployment(ctx context.Context, req *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.appsv1.Deployments(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}

func (c *Workload) ScaleDeployment(ctx context.Context, req *meta.ScaleRequest) (*v1.Scale, error) {
	return c.appsv1.Deployments(req.Scale.Namespace).UpdateScale(ctx, req.Scale.Name, req.Scale, req.Options)
}

// 原生并没有重新部署的功能, 通过变更注解时间来触发重新部署
// dpObj.Spec.Template.Annotations["cattle.io/timestamp"] = time.Now().Format(time.RFC3339)
func (c *Workload) ReDeploy(ctx context.Context, req *meta.GetRequest) (*appsv1.Deployment, error) {
	// 获取Deploy
	d, err := c.GetDeployment(ctx, req)
	if err != nil {
		return nil, err
	}
	// 添加一个时间戳来是Deploy对象发送变更
	d.Spec.Template.Annotations["mpaas/timestamp"] = time.Now().Format(time.RFC3339)
	return c.appsv1.Deployments(req.Namespace).Update(ctx, d, metav1.UpdateOptions{})
}

func (c *Workload) DeleteDeployment(ctx context.Context, req *meta.DeleteRequest) error {
	return c.appsv1.Deployments(req.Namespace).Delete(ctx, req.Name, req.Opts)
}
