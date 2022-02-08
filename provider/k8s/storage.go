package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListPersistentVolumeRequest struct {
	Namespace string
}

func (c *Client) ListPersistentVolume(ctx context.Context, req *ListPersistentVolumeRequest) (*v1.PersistentVolumeList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}

	return c.client.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
}

func (c *Client) ListPersistentVolumeClaims(ctx context.Context, req *ListPersistentVolumeRequest) (*v1.PersistentVolumeClaimList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().PersistentVolumeClaims(req.Namespace).List(ctx, metav1.ListOptions{})
}

func (c *Client) ListStorageClass(ctx context.Context, req *ListPersistentVolumeRequest) (*storagev1.StorageClassList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
}
