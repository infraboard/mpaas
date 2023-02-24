package network

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

func NewNetwork(cs *kubernetes.Clientset) *Client {
	return &Client{
		corev1:       cs.CoreV1(),
		networkingv1: cs.NetworkingV1(),
	}
}

type Client struct {
	corev1       corev1.CoreV1Interface
	networkingv1 networkingv1.NetworkingV1Interface
}
