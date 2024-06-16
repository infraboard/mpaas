package gateway

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListHttpRoute(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.HTTPRouteList, error) {
	d, err := c.dynamic.Resource(c.httpRouteResource()).Namespace("default").List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.HTTPRouteList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
