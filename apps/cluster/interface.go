package cluster

import (
	context "context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/pb/resource"
)

const (
	AppName = "clusters"
)

type Service interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*Cluster, error)
	RPCServer
}

func NewQueryClusterRequest() *QueryClusterRequest {
	return &QueryClusterRequest{
		Page:    request.NewDefaultPageRequest(),
		Filters: []*resource.LabelRequirement{},
		Label:   map[string]string{},
	}
}

func NewQueryClusterRequestFromHttp(r *restful.Request) *QueryClusterRequest {
	req := NewQueryClusterRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	req.Filters = policy.GetScopeFilterFromRequest(r)
	return req
}

func NewCreateClusterRequest() *CreateClusterRequest {
	return &CreateClusterRequest{
		Labels: map[string]string{},
	}
}

func NewDescribeClusterRequest(id string) *DescribeClusterRequest {
	return &DescribeClusterRequest{
		Id: id,
	}
}

func (r *DescribeClusterRequest) Validate() error {
	return nil
}

func NewDeleteClusterRequest(id string) *DeleteClusterRequest {
	return &DeleteClusterRequest{
		Id: id,
	}
}
