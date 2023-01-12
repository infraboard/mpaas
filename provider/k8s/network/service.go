package network

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Network) CreateService(ctx context.Context, req *v1.Service) (*v1.Service, error) {
	return c.corev1.Services(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Network) ListService(ctx context.Context, req *meta.ListRequest) (*v1.ServiceList, error) {
	return c.corev1.Services(req.Namespace).List(ctx, req.Opts)
}

func (c *Network) GetService(ctx context.Context, req *meta.GetRequest) (*v1.Service, error) {
	return c.corev1.Services(req.Namespace).Get(ctx, req.Name, req.Opts)
}
