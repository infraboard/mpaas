package gateway

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	"k8s.io/apimachinery/pkg/runtime"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (c *Client) ListGateway(
	ctx context.Context,
	req *meta.ListRequest) (
	*gatewayv1.GatewayList, error) {
	d, err := c.dynamic.Resource(c.gatewayResource()).Namespace("default").List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}

	list := new(gatewayv1.GatewayList)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(d.UnstructuredContent(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) CreateGateway(
	ctx context.Context,
	req *meta.CreateRequest) (
	*gatewayv1.GatewayList, error) {
	d, err := c.dynamic.Resource(c.gatewayResource()).Namespace("default").Create(ctx, nil, req.Opts)
	if err != nil {
		return nil, err
	}
	fmt.Println(d)
	return nil, nil
}
