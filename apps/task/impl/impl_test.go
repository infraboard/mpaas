package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl task.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(task.AppName).(task.Service)
}