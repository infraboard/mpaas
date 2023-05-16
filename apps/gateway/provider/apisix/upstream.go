package apisix

import (
	"context"
	"fmt"
)

// 创建Upstream 参考: https://apisix.apache.org/zh/docs/apisix/admin-api/#upstream
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

// 在 Upstream 中添加一个节点
func (c *Client) AddNodeToUpstream(ctx context.Context, in *AddNodeToUpstreamRequest) (*Upstream, error) {
	return nil, nil
}

type AddNodeToUpstreamRequest struct {
	UpStreamId string `json:"upstream_id"`
	*Node
}

// 更新 Upstream 中单个节点

func (c *Client) UpdateUpstreamNode(ctx context.Context, in *UpdateUpstreamNodeRequeset) (*Upstream, error) {
	return nil, nil
}

type UpdateUpstreamNodeRequeset struct {
	UpStreamId string `json:"upstream_id"`
	*Node
}

// 删除 Upstream 中的一个节点
func (c *Client) RemoveNodeFromUpstream(ctx context.Context, in *RemoveNodeFromUpstreamRequest) {

}

type RemoveNodeFromUpstreamRequest struct {
}
