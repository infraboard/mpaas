package k8s

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListIngressRequest struct {
	Namespace string
}

func (c *Client) ListIngress(ctx context.Context, req *ListServiceRequest) (*v1.IngressList, error) {
	if req.Namespace == "" {
		req.Namespace = apiv1.NamespaceDefault
	}
	return c.client.NetworkingV1().Ingresses(req.Namespace).List(ctx, metav1.ListOptions{})
}
