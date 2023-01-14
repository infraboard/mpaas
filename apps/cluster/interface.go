package cluster

import (
	context "context"
)

type Service interface {
	ClusterService
}

type ClusterService interface {
	RPCServer
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*Cluster, error)
}
