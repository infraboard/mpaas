package cluster

import (
	context "context"
	"strings"

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
		Page:       request.NewDefaultPageRequest(),
		Filters:    []*resource.LabelRequirement{},
		Label:      map[string]string{},
		ServiceIds: []string{},
		Names:      []string{},
		Ids:        []string{},
	}
}

func NewQueryClusterRequestFromHttp(r *restful.Request) *QueryClusterRequest {
	req := NewQueryClusterRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	req.Filters = policy.GetScopeFilterFromRequest(r)
	req.WithDeployment = r.QueryParameter("with_deploy") == "true"
	req.Filters = resource.ParseLabelRequirementListFromString(r.QueryParameter("filters"))
	req.AddServiceIds(r.QueryParameter("service_ids"))
	return req
}

func (req *QueryClusterRequest) AddServiceIds(ids string) {
	ids = strings.TrimSpace(ids)
	if ids == "" {
		return
	}

	idList := strings.Split(ids, ",")
	req.ServiceIds = append(req.ServiceIds, idList...)
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
