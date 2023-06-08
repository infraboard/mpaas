package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/clients/rest"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	client *rest.ClientSet
	ctx    = context.Background()
)

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Provider = "腾讯云"
	req.Region = "上海"
	req.Name = "生产环境(新)"

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

func TestQueryCluster(t *testing.T) {
	req := cluster.NewQueryClusterRequest()
	set, err := client.QueryCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeCluster(t *testing.T) {
	req := cluster.NewDescribeClusterRequest("cls-0f9m3dx3")
	ins, err := client.DescribeCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()
	conf := rest.NewDefaultConfig()
	client = rest.NewClient(conf)
}
