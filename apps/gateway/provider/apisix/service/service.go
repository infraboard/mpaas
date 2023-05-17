package service

import (
	"context"
	"fmt"
)

// 参考: https://apisix.apache.org/zh/docs/apisix/admin-api/#service

// 创建服务
// /apisix/admin/services
func (c *Client) CreateService(ctx context.Context, in *CreateServiceRequest) (
	*Service, error) {
	raw, err := c.c.
		Post("services").
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

// 查看服务列表
// /apisix/admin/services
func (c *Client) QueryService(ctx context.Context, in *QueryServiceRequest) (
	*ServiceList, error) {
	raw, err := c.c.
		Get("services").
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type QueryServiceRequest struct {
}

// 查看服务详情
// /apisix/admin/services/{id}
func (c *Client) DescribeService(ctx context.Context, in *DescribeServiceRequest) (
	*Service, error) {
	raw, err := c.c.
		Get("services").
		Suffix(in.ServiceId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DescribeServiceRequest struct {
	ServiceId string
}

// 删除服务
// /apisix/admin/services/{id}
func (c *Client) DeleteService(ctx context.Context, in *DeleteServiceRequest) (
	*Service, error) {
	raw, err := c.c.
		Delete("services").
		Suffix(in.ServiceId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DeleteServiceRequest struct {
	ServiceId string
}
