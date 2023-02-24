package storage

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
)

func (c *Client) ListPersistentVolume(ctx context.Context, req *meta.ListRequest) (*v1.PersistentVolumeList, error) {
	return c.corev1.PersistentVolumes().List(ctx, req.Opts)
}

func (c *Client) GetPersistentVolume(ctx context.Context, req *meta.GetRequest) (*v1.PersistentVolume, error) {
	return c.corev1.PersistentVolumes().Get(ctx, req.Name, req.Opts)
}
