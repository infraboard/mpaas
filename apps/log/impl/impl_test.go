package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/log"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl log.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(log.AppName).(log.Service)
}
