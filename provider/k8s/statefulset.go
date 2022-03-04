package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
)

func (c *Client) ListStatefulSet(ctx context.Context, req *ListRequest) (*appsv1.StatefulSetList, error) {
	return c.client.AppsV1().StatefulSets(req.Namespace).List(ctx, req.Opts)
}
