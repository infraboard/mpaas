package apisix

import (
	"github.com/infraboard/mcube/client/rest"
)

func NewClient(address, apiKey string) *Client {
	c := rest.NewRESTClient()
	c.SetBaseURL(address + "/apisix/admin")
	c.SetHeader("X-API-KEY", apiKey)
	return &Client{
		c: c,
	}
}

type Client struct {
	c *rest.RESTClient
}
