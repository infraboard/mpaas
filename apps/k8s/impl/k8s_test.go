package impl_test

import (
	"testing"

	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	ClusterId = "docker-desktop"
)

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	set, err := impl.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribeCluster(t *testing.T) {
	req := cluster.NewDescribeClusterRequest(ClusterId)
	ins, err := impl.DescribeCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Provider = "Docker"
	req.Region = "Local"
	req.Name = ClusterId

	req.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	ins, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateCluster(t *testing.T) {
	req := cluster.NewPatchClusterRequest(ClusterId)
	req.Spec.KubeConfig = tools.MustReadContentFile("test/kube_config.yml")
	ins, err := impl.UpdateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteCluster(t *testing.T) {
	req := cluster.NewDeleteClusterRequestWithID(ClusterId)
	ins, err := impl.DeleteCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
