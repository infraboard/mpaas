package k8s

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListEventRequest struct {
	Namespace string
}

func (c *Client) ListEvent(ctx context.Context, req *ListEventRequest) (*v1.EventList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.EventsV1().Events(req.Namespace).List(ctx, metav1.ListOptions{})
}
