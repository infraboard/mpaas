package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	req.Label["Env"] = "开发"
	req.WithService = true
	req.WithDeployment = true
	ds, err := impl.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Kind = cluster.KIND_MIDDLEWARE
	req.AccessConfig.Private.K8SCluster = "docker-desktop"
	req.AccessConfig.Private.K8SService = tools.MustReadContentFile("test/mongodb_service.yml")
	req.Labels["Env"] = "开发"
	req.Name = "MongoDB"
	ds, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestDescribeCluster(t *testing.T) {
	req := cluster.NewDescribeClusterRequest(conf.C.DEPLOY_CLUSTER_ID)
	ds, err := impl.DescribeCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}

func TestDeleteCluster(t *testing.T) {
	req := cluster.NewDeleteClusterRequest("2a90c4eec422c171")
	ds, err := impl.DeleteCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}
