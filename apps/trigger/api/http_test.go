package api_test

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/apps/trigger/api"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl *api.Handler
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetApi(trigger.AppName).(*api.Handler)
}
