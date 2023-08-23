package rpc_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"

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
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	c, err := rpc.NewClientSetFromEnv()
	if err != nil {
		panic(err)
	}
	client = c
}
