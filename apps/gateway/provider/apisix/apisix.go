package apisix

import (
	"github.com/infraboard/mcube/client/rest"
)

func DefaultClientConn() *rest.RESTClient {
	return NewClient(
		"http://127.0.0.1:9180",
		"edd1c9f034335f136f87ad84b625c8f1",
	).c
}

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
