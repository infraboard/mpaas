package apisix

import "context"

// 创建路由规则: https://apisix.apache.org/zh/docs/apisix/admin-api/#route
func (c *Client) CreateRoute(ctx context.Context, in *CreateRouteRequest) (
	*Route, error) {
	return nil, nil
}
