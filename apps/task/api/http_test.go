package api_test

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/task/api"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl *api.Handler
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetApi(task.AppName).(*api.Handler)
}
