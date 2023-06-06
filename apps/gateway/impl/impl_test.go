package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/gateway"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl gateway.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(gateway.AppName).(gateway.Service)
}
