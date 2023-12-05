package upstream

import "github.com/infraboard/mcube/v2/tools/pretty"

func NewDescribeUpstreamRequest(upstreamId string) *DescribeUpstreamRequest {
	return &DescribeUpstreamRequest{
		UpStreamId: upstreamId,
	}
}

type DescribeUpstreamRequest struct {
	UpStreamId string `json:"upstream_id"`
}

type AddNodeToUpstreamRequest struct {
	UpStreamId string  `json:"upstream_id"`
	Nodes      []*Node `json:"nodes"`
}

func (r *AddNodeToUpstreamRequest) ToJSON() string {
	return pretty.ToJSON(r)
}

type UpdateUpstreamNodeRequeset struct {
	UpStreamId string `json:"upstream_id"`
	*Node
}

type RemoveNodeFromUpstreamRequest struct {
}

type DeleteUpstreamRequest struct {
	UpStreamId string `json:"upstream_id"`
}
