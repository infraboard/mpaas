package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListConfigMap(ctx context.Context, req *ListRequest) (*v1.ConfigMapList, error) {
	return c.client.CoreV1().ConfigMaps(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) CreateConfigMap(ctx context.Context, req *v1.ConfigMap) (*v1.ConfigMap, error) {
	return c.client.CoreV1().ConfigMaps(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}
