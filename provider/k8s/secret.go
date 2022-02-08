package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListSecretRequest struct {
	Namespace string
}

func (c *Client) ListSecret(ctx context.Context, req *ListServiceRequest) (*v1.SecretList, error) {
	if req.Namespace == "" {
		req.Namespace = v1.NamespaceDefault
	}
	return c.client.CoreV1().Secrets(req.Namespace).List(ctx, metav1.ListOptions{})
}
