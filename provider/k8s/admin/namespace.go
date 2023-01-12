package admin

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Admin) ListNamespace(ctx context.Context, req *meta.ListRequest) (*v1.NamespaceList, error) {
	set, err := c.corev1.Namespaces().List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}
	if req.SkipManagedFields {
		for i := range set.Items {
			set.Items[i].ManagedFields = nil
		}
	}
	return set, nil
}

func (c *Admin) CreateNamespace(ctx context.Context, req *v1.Namespace) (*v1.Namespace, error) {
	return c.corev1.Namespaces().Create(ctx, req, metav1.CreateOptions{})
}

func (c *Admin) ListResourceQuota(ctx context.Context) (*v1.ResourceQuotaList, error) {
	return c.corev1.ResourceQuotas("").List(ctx, metav1.ListOptions{})
}

func (c *Admin) CreateResourceQuota(ctx context.Context, req *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	return c.corev1.ResourceQuotas(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}
