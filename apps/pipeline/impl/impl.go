package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	pipeline.UnimplementedRPCServer
	ioc.IocObjectImpl
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return pipeline.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	pipeline.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
