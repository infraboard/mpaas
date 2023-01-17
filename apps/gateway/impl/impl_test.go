package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/gateway"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl gateway.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(gateway.AppName).(gateway.Service)
}
