package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	cluster.UnimplementedRPCServer

	mcenter *rpc.ClientSet
	cluster cluster.Service
	ioc.IocObjectImpl
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	i.mcenter = rpc.C()
	i.cluster = ioc.GetController(cluster.AppName).(cluster.Service)
	return nil
}

func (i *impl) Name() string {
	return cluster.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	cluster.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
