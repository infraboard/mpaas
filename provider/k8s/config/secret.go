package config

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Config) CreateSecret(ctx context.Context, req *v1.Secret) (*v1.Secret, error) {
	return c.corev1.Secrets(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Config) ListSecret(ctx context.Context, req *meta.ListRequest) (*v1.SecretList, error) {
	return c.corev1.Secrets(req.Namespace).List(ctx, req.Opts)
}

func (c *Config) GetSecret(ctx context.Context, req *meta.GetRequest) (*v1.Secret, error) {
	return c.corev1.Secrets(req.Namespace).Get(ctx, req.Name, req.Opts)
}

func (c *Config) UpdateSecret(ctx context.Context, req *v1.Secret) (*v1.Secret, error) {
	return c.corev1.Secrets(req.Namespace).Update(ctx, req, metav1.UpdateOptions{})
}

func (c *Config) FindOrCreate(ctx context.Context, secret *v1.Secret) error {
	req := meta.NewGetRequest(secret.Name).WithNamespace(secret.Namespace)
	_, err := c.GetSecret(ctx, req)
	if errors.IsNotFound(err) {
		s, err := c.CreateSecret(ctx, secret)
		if err != nil {
			return err
		}
		// 返回创建的值
		*secret = *s
	}
	return nil
}
