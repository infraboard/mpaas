package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl deploy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(deploy.AppName).(deploy.Service)
}
