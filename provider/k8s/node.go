package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListNode(ctx context.Context) (*v1.NodeList, error) {
	return c.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
}
