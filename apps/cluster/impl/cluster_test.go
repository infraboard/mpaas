package impl_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	set, err := impl.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeCluster(t *testing.T) {
	req := cluster.NewDescribeClusterRequest("k8s-test")
	set, err := impl.DescribeCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Vendor = "腾讯云"
	req.Region = "上海"
	req.Name = "k8s-test"

	req.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	ins, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateCluster(t *testing.T) {
	req := cluster.NewPatchClusterRequest("k8s-test")
	req.Spec.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	set, err := impl.UpdateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDeleteCluster(t *testing.T) {
	req := cluster.NewDeleteClusterRequestWithID("k8s-test")
	set, err := impl.DeleteCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
