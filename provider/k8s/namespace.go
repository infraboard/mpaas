package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListNamespace(ctx context.Context) (*v1.NamespaceList, error) {
	return c.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
}

func (c *Client) CreateNamespace(ctx context.Context, req *v1.Namespace) (*v1.Namespace, error) {
	return c.client.CoreV1().Namespaces().Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListResourceQuota(ctx context.Context) (*v1.ResourceQuotaList, error) {
	return c.client.CoreV1().ResourceQuotas("").List(ctx, metav1.ListOptions{})
}

func (c *Client) CreateResourceQuota(ctx context.Context, req *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	return c.client.CoreV1().ResourceQuotas(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}
