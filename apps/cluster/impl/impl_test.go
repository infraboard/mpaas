package impl

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl cluster.Service
	ctx  = context.Background()
)

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Vendor = "腾讯云"
	req.Region = "上海"
	req.Name = "生产环境"

	kubeConf, err := tools.ReadFile("test/kube_config.yml")
	if err != nil {
		t.Fatal(err)
	}
	req.KubeConfig = string(kubeConf)
	ins, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
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

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(cluster.AppName).(cluster.Service)
}
