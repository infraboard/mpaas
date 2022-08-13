package k8s

import (
	"context"
	"net/http"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
)

func NewListRequestFromHttp(r *http.Request) *ListRequest {
	qs := r.URL.Query()

	req := &ListRequest{
		Namespace: qs.Get("namespace"),
	}

	return req
}

func (c *Client) ListDeployment(ctx context.Context, req *ListRequest) (*appsv1.DeploymentList, error) {
	return c.client.AppsV1().Deployments(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetDeployment(ctx context.Context, req *GetRequest) (*appsv1.Deployment, error) {
	return c.client.AppsV1().Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
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

func NewUpdateScaleRequest() *UpdateScaleRequest {
	return &UpdateScaleRequest{
		Scale:   &v1.Scale{},
		Options: metav1.UpdateOptions{},
	}
}

type UpdateScaleRequest struct {
	Scale   *v1.Scale
	Options metav1.UpdateOptions
}

func (c *Client) UpdateDeploymentScale(ctx context.Context, req *UpdateScaleRequest) (*v1.Scale, error) {
	return c.client.AppsV1().Deployments(req.Scale.Namespace).UpdateScale(ctx, req.Scale.Name, req.Scale, req.Options)
}

type ReDeployRequest struct {
	Namespace string
	Name      string
}

// 原生并没有重新部署的功能, 通过变更注解时间来触发重新部署
// dpObj.Spec.Template.Annotations["cattle.io/timestamp"] = time.Now().Format(time.RFC3339)
func (c *Client) ReDeploy(ctx context.Context, req *GetRequest) (*appsv1.Deployment, error) {
	// 获取Deploy
	d, err := c.GetDeployment(ctx, req)
	if err != nil {
		return nil, err
	}
	// 添加一个时间戳来是Deploy对象发送变更
	d.Spec.Template.Annotations["devcloud/timestamp"] = time.Now().Format(time.RFC3339)
	return c.client.AppsV1().Deployments(req.Namespace).Update(ctx, d, metav1.UpdateOptions{})
}

func (c *Client) DeleteDeployment(ctx context.Context, req *DeleteRequest) error {
	return c.client.AppsV1().Deployments(req.Namespace).Delete(ctx, req.Name, req.Opts)
}
