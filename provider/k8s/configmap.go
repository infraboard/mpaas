package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListConfigMapRequest struct {
	Namespace string
}

func (c *Client) ListConfigMap(ctx context.Context, req *ListServiceRequest) (*v1.ConfigMapList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().ConfigMaps(req.Namespace).List(ctx, metav1.ListOptions{})
}
