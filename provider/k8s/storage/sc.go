package storage

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"

	"github.com/infraboard/mpaas/provider/k8s/meta"
)

func (c *Client) ListStorageClass(ctx context.Context, req *meta.ListRequest) (*storagev1.StorageClassList, error) {
	return c.storagev1.StorageClasses().List(ctx, req.Opts)
}

func (c *Client) GetStorageClass(ctx context.Context, req *meta.GetRequest) (*storagev1.StorageClass, error) {
	return c.storagev1.StorageClasses().Get(ctx, req.Name, req.Opts)
}
