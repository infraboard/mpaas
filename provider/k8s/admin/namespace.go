package admin

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListNamespace(ctx context.Context, req *meta.ListRequest) (*v1.NamespaceList, error) {
	return c.corev1.Namespaces().List(ctx, req.Opts)
}

func (c *Client) GetNamespace(ctx context.Context, req *meta.GetRequest) (*v1.Namespace, error) {
	return c.corev1.Namespaces().Get(ctx, req.Name, req.Opts)
}

func (c *Client) CreateNamespace(ctx context.Context, req *v1.Namespace) (*v1.Namespace, error) {
	return c.corev1.Namespaces().Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListResourceQuota(ctx context.Context) (*v1.ResourceQuotaList, error) {
	return c.corev1.ResourceQuotas("").List(ctx, metav1.ListOptions{})
}

func (c *Client) CreateResourceQuota(ctx context.Context, req *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	return c.corev1.ResourceQuotas(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}
