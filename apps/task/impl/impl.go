package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/conf"

	// 加载并初始化Runner
	"github.com/infraboard/mpaas/apps/task/runner"
	_ "github.com/infraboard/mpaas/apps/task/runner/k8s"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	task.UnimplementedRPCServer

	job job.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	i.job = app.GetInternalApp(job.AppName).(job.Service)
	return runner.Init()
}

func (i *impl) Name() string {
	return task.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	task.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
