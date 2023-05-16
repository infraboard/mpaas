package apisix

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/tools/pretty"
)

// 参考: https://apisix.apache.org/zh/docs/apisix/admin-api/#upstream

// 创建Upstream
// /apisix/admin/upstreams
func (c *Client) CreateUpstream(ctx context.Context, in *CreateUpstreamRequeset) (
	*Upstream, error) {
	raw, err := c.c.
		Post("upstreams").
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

// 查询Upstream详情
// /apisix/admin/upstreams/{id}
func (c *Client) DescribeUpstream(ctx context.Context, in *DescribeUpstreamRequest) (
	*Upstream, error) {
	raw, err := c.c.
		Get("upstreams").
		Suffix(in.UpStreamId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DescribeUpstreamRequest struct {
	UpStreamId string `json:"upstream_id"`
}

// 在 Upstream 中添加一个节点
// /apisix/admin/upstreams/100
func (c *Client) AddNodeToUpstream(ctx context.Context, in *AddNodeToUpstreamRequest) (
	*Upstream, error) {
	raw, err := c.c.
		Patch("upstreams").
		Suffix(in.UpStreamId).
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type AddNodeToUpstreamRequest struct {
	UpStreamId string  `json:"upstream_id"`
	Nodes      []*Node `json:"nodes"`
}

func (r *AddNodeToUpstreamRequest) ToJSON() string {
	return pretty.ToJSON(r)
}

// 更新 Upstream 中单个节点
func (c *Client) UpdateUpstreamNode(ctx context.Context, in *UpdateUpstreamNodeRequeset) (
	*Upstream, error) {
	return nil, nil
}

type UpdateUpstreamNodeRequeset struct {
	UpStreamId string `json:"upstream_id"`
	*Node
}

// 删除 Upstream 中的一个节点
func (c *Client) RemoveNodeFromUpstream(ctx context.Context, in *RemoveNodeFromUpstreamRequest) (
	*Upstream, error) {
	return nil, nil
}

type RemoveNodeFromUpstreamRequest struct {
}

// 删除Upstream
// /apisix/admin/upstreams/{id}
func (c *Client) DeleteUpstream(ctx context.Context, in *DeleteUpstreamRequest) (
	*Upstream, error) {
	raw, err := c.c.
		Delete("upstreams").
		Suffix(in.UpStreamId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

type DeleteUpstreamRequest struct {
	UpStreamId string `json:"upstream_id"`
}
