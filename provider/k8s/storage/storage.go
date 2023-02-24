package storage

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

func NewStorage(cs *kubernetes.Clientset) *Client {
	return &Client{
		corev1:    cs.CoreV1(),
		storagev1: cs.StorageV1(),
	}
}

type Client struct {
	corev1    corev1.CoreV1Interface
	storagev1 storagev1.StorageV1Interface
}
