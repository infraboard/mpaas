package upstream_test

import (
	"context"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/upstream"
)

var (
	client *upstream.Client
	ctx    = context.Background()
)

func init() {
	zap.DevelopmentSetup()
	client = upstream.NewClient(apisix.DefaultClientConn())
}
