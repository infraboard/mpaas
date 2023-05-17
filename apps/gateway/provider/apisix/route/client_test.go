package route_test

import (
	"context"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix"
	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/route"
)

var (
	client route.Service
	ctx    = context.Background()
)

func init() {
	zap.DevelopmentSetup()
	client = route.NewClient(apisix.DefaultClientConn())
}
