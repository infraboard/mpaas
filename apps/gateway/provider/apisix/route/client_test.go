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
	conn := apisix.NewClient(
		"http://127.0.0.1:9180",
		"edd1c9f034335f136f87ad84b625c8f1",
	).Conn()
	client = route.NewClient(conn)
}
