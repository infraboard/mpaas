package storage

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Storage) ListPersistentVolumeClaims(ctx context.Context, req *meta.ListRequest) (*v1.PersistentVolumeClaimList, error) {
	return c.corev1.PersistentVolumeClaims(req.Namespace).List(ctx, req.Opts)
}

func (c *Storage) GetPersistentVolumeClaims(ctx context.Context, req *meta.GetRequest) (*v1.PersistentVolumeClaim, error) {
	return c.corev1.PersistentVolumeClaims(req.Namespace).Get(ctx, req.Name, req.Opts)
}
