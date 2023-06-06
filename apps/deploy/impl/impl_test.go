package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl deploy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(deploy.AppName).(deploy.Service)
}
