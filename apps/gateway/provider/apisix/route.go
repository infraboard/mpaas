package apisix

import (
	"context"
	"fmt"
)

// 路由规则: https://apisix.apache.org/zh/docs/apisix/admin-api/#route

// 创建路由规则
// /apisix/admin/routes
func (c *Client) CreateRoute(ctx context.Context, in *CreateRouteRequest) (
	*Route, error) {
	raw, err := c.c.
		Post("routes").
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

// 查询路由列表
// /apisix/admin/routes
func (c *Client) QueryRoute(ctx context.Context, in *QueryRouteRequest) (
	*RouteList, error) {
	raw, err := c.c.
		Get("routes").
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type QueryRouteRequest struct {
}

// 查询路由详情
// /apisix/admin/routes/{id}
func (c *Client) DescribeRoute(ctx context.Context, in *DescribeRouteRequest) (
	*Route, error) {
	raw, err := c.c.
		Get("routes").
		Suffix(in.RouteId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DescribeRouteRequest struct {
	RouteId string
}

// 更新路由
// /apisix/admin/routes/{id}
func (c *Client) UpdateRoute(ctx context.Context, in *UpdateRouteRequest) (
	*Route, error) {
	return nil, nil
}

type UpdateRouteRequest struct {
}

// 删除路由
// /apisix/admin/routes/{id}
func (c *Client) DeleteRoute(ctx context.Context, in *DeleteRouteRequest) (
	*Route, error) {
	raw, err := c.c.
		Delete("routes").
		Suffix(in.RouteId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DeleteRouteRequest struct {
	RouteId string
}
