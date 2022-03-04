package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

type ListPersistentVolumeRequest struct {
	Namespace string
}

func (c *Client) ListPersistentVolume(ctx context.Context, req *ListRequest) (*v1.PersistentVolumeList, error) {
	return c.client.CoreV1().PersistentVolumes().List(ctx, req.Opts)
}

func (c *Client) ListPersistentVolumeClaims(ctx context.Context, req *ListRequest) (*v1.PersistentVolumeClaimList, error) {
	return c.client.CoreV1().PersistentVolumeClaims(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) ListStorageClass(ctx context.Context, req *ListRequest) (*storagev1.StorageClassList, error) {
	return c.client.StorageV1().StorageClasses().List(ctx, req.Opts)
}
