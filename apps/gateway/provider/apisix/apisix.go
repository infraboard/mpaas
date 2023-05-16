package apisix

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/client/rest"
)

func NewClient(address, apiKey string) *Client {
	c := rest.NewRESTClient()
	c.SetHeader(restful.HEADER_ContentType, restful.MIME_JSON)
	c.SetHeader(restful.HEADER_Accept, restful.MIME_JSON)
	c.SetBaseURL(address + "/apisix/admin")
	c.SetHeader("X-API-KEY", apiKey)
	return &Client{
		c: c,
	}
}

type Client struct {
	c *rest.RESTClient
}
