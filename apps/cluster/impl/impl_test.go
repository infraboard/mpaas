package impl_test

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl cluster.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(cluster.AppName).(cluster.Service)
}
