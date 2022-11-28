package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateStatefulSet(ctx context.Context, req *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	return c.client.AppsV1().StatefulSets(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListStatefulSet(ctx context.Context, req *ListRequest) (*appsv1.StatefulSetList, error) {
	return c.client.AppsV1().StatefulSets(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetStatefulSet(ctx context.Context, req *GetRequest) (*appsv1.StatefulSet, error) {
	return c.client.AppsV1().StatefulSets(req.Namespace).Get(ctx, req.Name, req.Opts)
}
