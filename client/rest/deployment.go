package rest

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/restful/accessor/yamlk8s"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
)

func (c *ClientSet) CreateDeployment(ctx context.Context, req *appsv1.Deployment) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.Group("clusters").
		Group(req.Annotations[cluster.ANNOTATION_CLUSTER_ID]).
		Post("deployments").
		Body(req).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *ClientSet) CreateDeploymentByYaml(ctx context.Context, clusterName, yamlString string) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.Group("clusters").
		Group(clusterName).
		Post("deployments").
		Header(restful.HEADER_ContentType, yamlk8s.MIME_YAML).
		Body(yamlString).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *ClientSet) QueryDeployment(ctx context.Context, req *meta.ListRequest) (
	*cluster.ClusterSet, error) {
	set := cluster.NewClusterSet()

	err := c.c.Group("clusters").
		Group(req.Namespace).
		Get("deployments").
		Body(req).
		Do(ctx).
		Into(response.NewData(set))
	if err != nil {
		return nil, err
	}

	return set, nil
}
