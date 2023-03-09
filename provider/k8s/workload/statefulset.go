package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateStatefulSet(ctx context.Context, req *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	return c.appsv1.StatefulSets(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListStatefulSet(ctx context.Context, req *meta.ListRequest) (*appsv1.StatefulSetList, error) {
	return c.appsv1.StatefulSets(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetStatefulSet(ctx context.Context, req *meta.GetRequest) (*appsv1.StatefulSet, error) {
	return c.appsv1.StatefulSets(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Client) DeleteStatefulSet(ctx context.Context, req *meta.DeleteRequest) error {
	return c.appsv1.StatefulSets(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

func GetStatefulSetStatus(*appsv1.StatefulSet) *WorkloadStatus {
	return nil
}
