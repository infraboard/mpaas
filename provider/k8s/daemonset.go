package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListDaemonSet(ctx context.Context, req *ListDeploymentRequest) (*appsv1.DaemonSetList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.AppsV1().DaemonSets(req.Namespace).List(ctx, metav1.ListOptions{})
}
