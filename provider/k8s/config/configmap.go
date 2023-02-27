package config

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListConfigMap(ctx context.Context, req *meta.ListRequest) (*v1.ConfigMapList, error) {
	return c.corev1.ConfigMaps(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetConfigMap(ctx context.Context, req *meta.GetRequest) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Client) DeleteConfigMap(ctx context.Context, req *meta.DeleteRequest) error {
	return c.corev1.ConfigMaps(req.Namespace).Delete(ctx, req.Name, req.Opts)
}

func (c *Client) CreateConfigMap(ctx context.Context, req *v1.ConfigMap) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) UpdateConfigMap(ctx context.Context, req *v1.ConfigMap) (*v1.ConfigMap, error) {
	return c.corev1.ConfigMaps(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}

func (c *Client) FindOrCreateConfigMap(ctx context.Context, cm *v1.ConfigMap) error {
	req := meta.NewGetRequest(cm.Name).WithNamespace(cm.Namespace)
	_, err := c.GetSecret(ctx, req)
	if errors.IsNotFound(err) {
		s, err := c.CreateConfigMap(ctx, cm)
		if err != nil {
			return err
		}
		// 返回创建的值
		*cm = *s
	}
	return nil
}
