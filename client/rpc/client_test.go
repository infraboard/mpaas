package rpc_test

import (
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/client/rpc"
	"github.com/infraboard/mpaas/conf"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	c, err := rpc.NewClientSet(conf.C().Mcenter)
	if should.NoError(err) {
		fmt.Println(c)
	}
}

func init() {
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}
