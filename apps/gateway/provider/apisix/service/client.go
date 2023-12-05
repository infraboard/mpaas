package service

import "github.com/infraboard/mcube/v2/client/rest"

func NewClient(c *rest.RESTClient) *Client {
	return &Client{
		c: c,
	}
}

// 路由规则: https://apisix.apache.org/zh/docs/apisix/admin-api/#route
type Client struct {
	c *rest.RESTClient
}
