package config

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Config) ListConfigMap(ctx context.Context, req *meta.ListRequest) (*v1.ConfigMapList, error) {
	return c.corev1.ConfigMaps(req.Namespace).List(ctx, req.Opts)
}

func (c *Config) GetConfigMap(ctx context.Context, req *meta.GetRequest) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Config) CreateConfigMap(ctx context.Context, req *v1.ConfigMap) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Config) UpdateConfigMap(ctx context.Context, req *v1.ConfigMap) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}
