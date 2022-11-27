package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/client/rest"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	client *rest.ClientSet
	ctx    = context.Background()
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
	ins, err := client.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	conf := rest.NewDefaultConfig()
	client = rest.NewClient(conf)
}
