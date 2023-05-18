package apisix

import (
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/route"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/service"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/upstream"
)

func NewClient(address, apiKey string) *ClientSet {
	c := rest.NewRESTClient()
	c.SetBaseURL(address + "/apisix/admin")
	c.SetHeader("X-API-KEY", apiKey)
	c.EnableTrace()
	return &ClientSet{
		c: c,
	}
}

type ClientSet struct {
	c *rest.RESTClient
}

func (c *ClientSet) Conn() *rest.RESTClient {
	return c.c
}

func (c *ClientSet) Upstream() *upstream.Client {
	return upstream.NewClient(c.c)
}

func (c *ClientSet) Route() *route.Client {
	return route.NewClient(c.c)
}

func (c *ClientSet) Service() *service.Client {
	return service.NewClient(c.c)
}
