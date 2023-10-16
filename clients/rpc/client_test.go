package rpc_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/clients/rpc"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestQueryDeployment(t *testing.T) {
	req := deploy.NewQueryDeploymentRequest()
	set, err := client.Deploy().QueryDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	err := ioc.ConfigIocObject(ioc.NewLoadConfigRequest())
	if err != nil {
		panic(err)
	}
	client = rpc.C()
}
