package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateSecret(ctx context.Context, req *v1.Secret) (*v1.Secret, error) {
	return c.client.CoreV1().Secrets(req.Namespace).Create(ctx, req, metav1.CreateOptions{})
}

func (c *Client) ListSecret(ctx context.Context, req *ListRequest) (*v1.SecretList, error) {
	return c.client.CoreV1().Secrets(req.Namespace).List(ctx, req.Opts)
}

func (c *Client) GetSecret(ctx context.Context, req *GetRequest) (*v1.Secret, error) {
	return c.client.CoreV1().Secrets(req.Namespace).Get(ctx, req.Name, req.Opts)
}
