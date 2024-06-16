package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListGatewayClass(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.GatewayClass, error) {
	d, err := c.dynamic.Resource(c.gatewayClassResource()).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.GatewayClass)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
