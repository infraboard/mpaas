package config

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func NewConfig(cs *kubernetes.Clientset) *Config {
	return &Config{
		corev1: cs.CoreV1(),
	}
}

type Config struct {
	corev1 corev1.CoreV1Interface
}
