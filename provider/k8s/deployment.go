package k8s

import (
	"context"
	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
)

func NewListDeploymentRequestFromHttp(r *http.Request) *ListDeploymentRequest {
	qs := r.URL.Query()

	req := &ListDeploymentRequest{
		Namespace: qs.Get("namespace"),
	}

	return req
}

type ListDeploymentRequest struct {
	Namespace string
}

func (c *Client) ListDeployment(ctx context.Context, req *ListDeploymentRequest) (*appsv1.DeploymentList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.AppsV1().Deployments(req.Namespace).List(ctx, metav1.ListOptions{})
}

func (c *Client) WatchDeployment(ctx context.Context, req *appsv1.Deployment) (watch.Interface, error) {
	return c.client.AppsV1().Deployments(req.Namespace).Watch(ctx, metav1.ListOptions{})
}

func (c *Client) CreateDeployment(ctx context.Context, req *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.client.AppsV1().Deployments(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) UpdateDeployment(ctx context.Context, req *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.client.AppsV1().Deployments(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}

func (c *Client) UpdateScale(ctx context.Context) {
	c.client.AppsV1().Deployments("").UpdateScale(ctx, "", nil, metav1.UpdateOptions{})
}

// 原生并没有重新部署的功能, 通过变更注解时间来触发重新部署
// dpObj.Spec.Template.Annotations["cattle.io/timestamp"] = time.Now().Format(time.RFC3339)
func (c *Client) ReDeploy(ctx context.Context) {

}
