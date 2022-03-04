package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
)

func (c *Client) ListDaemonSet(ctx context.Context, req *ListRequest) (*appsv1.DaemonSetList, error) {
	return c.client.AppsV1().DaemonSets(req.Namespace).List(ctx, req.Opts)
}
