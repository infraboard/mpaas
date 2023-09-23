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
	req.ServiceId = conf.C.MCENTER_SERVICE_ID
	req.Labels["Env"] = "开发"
	req.Name = "默认集群"
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
	req := cluster.NewDeleteClusterRequest("xxx")
	ds, err := impl.DeleteCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ds))
}
