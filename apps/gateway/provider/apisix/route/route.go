package route

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/common"
)

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
	list := NewRouteList()
	resp := common.NewReponseList()
	err := c.c.
		Get("routes").
		Do(ctx).
		Into(resp)

	resp.Values(list)
	return list, err
}

// 查询路由详情
// /apisix/admin/routes/{id}
func (c *Client) DescribeRoute(ctx context.Context, in *DescribeRouteRequest) (
	*Route, error) {
	resp := common.NewReponse()
	err := c.c.
		Get("routes").
		Suffix(in.RouteId).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	r := NewRoute()
	err = resp.GetValue(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 更新路由
// /apisix/admin/routes/{id}
func (c *Client) UpdateRoute(ctx context.Context, in *UpdateRouteRequest) (
	*Route, error) {
	return nil, nil
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
