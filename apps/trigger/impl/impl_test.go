package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl trigger.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(trigger.AppName).(trigger.Service)
}
