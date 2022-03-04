package k8s

import (
	"context"

	v1 "k8s.io/api/networking/v1"
)

func (c *Client) ListIngress(ctx context.Context, req *ListRequest) (*v1.IngressList, error) {
	return c.client.NetworkingV1().Ingresses(req.Namespace).List(ctx, req.Opts)
}
