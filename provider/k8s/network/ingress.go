package network

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/networking/v1"
)

func (c *Access) ListIngress(ctx context.Context, req *meta.ListRequest) (*v1.IngressList, error) {
	return c.networkingv1.Ingresses(req.Namespace).List(ctx, req.Opts)
}
