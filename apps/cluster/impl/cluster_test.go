package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/tools"
)

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Vendor = "腾讯云"
	req.Region = "上海"
	req.Name = "生产环境"

	req.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	ins, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateCluster(t *testing.T) {
	req := cluster.NewPatchClusterRequest("cls-ot6msuhj-local")
	req.Data.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	set, err := impl.UpdateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	set, err := impl.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDeleteCluster(t *testing.T) {
	req := cluster.NewDeleteClusterRequestWithID("cls-ot6msuhj-local")
	set, err := impl.DeleteCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
