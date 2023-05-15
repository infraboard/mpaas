package apisix

import "context"

// 创建服务, 参考: https://apisix.apache.org/zh/docs/apisix/admin-api/#service
func (c *Client) CreateService(ctx context.Context, in *CreateRouteRequest) (
	*Service, error) {
	return nil, nil
}
