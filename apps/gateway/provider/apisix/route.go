package apisix

import "context"

// 创建路由规则: https://apisix.apache.org/zh/docs/apisix/admin-api/#route
func (c *Client) CreateRoute(ctx context.Context, in *CreateRouteRequest) (
	*Route, error) {
	return nil, nil
}

func (c *Client) UpdateRoute(ctx context.Context, in *UpdateRouteRequest) (
	*Route, error) {
	return nil, nil
}

type UpdateRouteRequest struct {
}

// 删除路由
func (c *Client) DeleteRoute(ctx context.Context, in *DeleteRouteRequest) (
	*Route, error) {
	return nil, nil
}

type DeleteRouteRequest struct {
}
