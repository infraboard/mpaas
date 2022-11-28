package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateService(ctx context.Context, req *v1.Service) (*v1.Service, error) {
	return c.client.CoreV1().Services(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListService(ctx context.Context, req *ListRequest) (*v1.ServiceList, error) {
	return c.client.CoreV1().Services(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetService(ctx context.Context, req *GetRequest) (*v1.Service, error) {
	return c.client.CoreV1().Services(req.Namespace).Get(ctx, req.Name, req.Opts)
}
