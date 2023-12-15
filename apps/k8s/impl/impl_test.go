package impl_test

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl cluster.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(cluster.AppName).(cluster.Service)
}
