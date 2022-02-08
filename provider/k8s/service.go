package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListServiceRequest struct {
	Namespace string
}

func (c *Client) ListService(ctx context.Context, req *ListServiceRequest) (*v1.ServiceList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().Services(req.Namespace).List(ctx, metav1.ListOptions{})
}

func (c *Client) CreateService(ctx context.Context, req *v1.Service) (*v1.Service, error) {
	return c.client.CoreV1().Services(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}
