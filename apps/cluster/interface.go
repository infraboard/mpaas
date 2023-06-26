package cluster

import (
	context "context"

	"github.com/infraboard/mcube/http/request"
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
		Page: request.NewDefaultPageRequest(),
	}
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
