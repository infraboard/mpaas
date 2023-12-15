package impl_test

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mpaas/apps/audit"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl audit.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(audit.AppName).(audit.Service)
}
