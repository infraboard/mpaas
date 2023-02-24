package admin

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Client) ListNode(ctx context.Context, req *meta.ListRequest) (*v1.NodeList, error) {
	return c.corev1.Nodes().List(ctx, req.Opts)
}

func (c *Client) GetNode(ctx context.Context, req *meta.GetRequest) (*v1.Node, error) {
	return c.corev1.Nodes().Get(ctx, req.Name, req.Opts)
}
