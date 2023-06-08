package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	trigger.UnimplementedRPCServer
	ioc.IocObjectImpl

	// 构建配置
	build build.Service
	// 执行流水线
	task task.PipelineService
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	i.build = ioc.GetController(build.AppName).(build.Service)
	i.task = ioc.GetController(task.AppName).(task.Service)
	return nil
}

func (i *impl) Name() string {
	return trigger.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	trigger.RegisterRPCServer(server, i)
}
