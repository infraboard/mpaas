package gateway

import (
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func NewGateway(restconf *rest.Config, resources *meta.ApiResourceList) *Client {
	return &Client{
		resources: resources,
		dynamic:   dynamic.NewForConfigOrDie(restconf),
	}
}

// https://github.com/kubernetes-sigs/gateway-api
type Client struct {
	resources *meta.ApiResourceList
	dynamic   *dynamic.DynamicClient
}

func (c *Client) resource(resourceName string) schema.GroupVersionResource {
	r := c.resources.GetResourceByName(resourceName)
	return schema.GroupVersionResource{
		Group:    r.Group,
		Version:  r.Version,
		Resource: resourceName,
	}
}

func (c *Client) gatewayResource() schema.GroupVersionResource {
	return c.resource("gateways")
}

func (c *Client) gatewayClassResource() schema.GroupVersionResource {
	return c.resource("gatewayclasses")
}

func (c *Client) grpcRouteResource() schema.GroupVersionResource {
	return c.resource("grpcroutes")
}

func (c *Client) httpRouteResource() schema.GroupVersionResource {
	return c.resource("httproutes")
}

func (c *Client) referenceGrantResource() schema.GroupVersionResource {
	r := c.resource("referencegrants")
	r.Version = "v1beta1"
	return r
}
