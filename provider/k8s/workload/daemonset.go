package workload

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
)

func (c *Workload) ListDaemonSet(ctx context.Context, req *meta.ListRequest) (*appsv1.DaemonSetList, error) {
	return c.appsv1.DaemonSets(req.Namespace).List(ctx, req.Opts)
}
