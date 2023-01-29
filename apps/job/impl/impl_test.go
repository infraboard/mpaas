package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl job.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(job.AppName).(job.Service)
}
