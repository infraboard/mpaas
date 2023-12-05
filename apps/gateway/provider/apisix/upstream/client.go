package upstream

import "github.com/infraboard/mcube/v2/client/rest"

func NewClient(c *rest.RESTClient) *Client {
	return &Client{
		c: c,
	}
}

// 参考: https://apisix.apache.org/zh/docs/apisix/admin-api/#upstream
type Client struct {
	c *rest.RESTClient
}
