package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl approval.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(approval.AppName).(approval.Service)
}
