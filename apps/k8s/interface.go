package k8s

import (
	context "context"
)

const (
	AppName = "k8s_clusters"
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
