package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
)

func (c *Client) ListNode(ctx context.Context, req *ListRequest) (*v1.NodeList, error) {
	return c.client.CoreV1().Nodes().List(ctx, req.Opts)
}
