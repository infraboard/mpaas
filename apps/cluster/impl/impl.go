package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	cluster.UnimplementedRPCServer
	ioc.ObjectImpl

	k8s     k8s.Service
	mcenter *rpc.ClientSet
	deploy  deploy.RPCServer
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	i.mcenter = rpc.C()
	i.deploy = ioc.GetController(deploy.AppName).(deploy.Service)
	i.k8s = ioc.GetController(k8s.AppName).(k8s.Service)
	return nil
}

func (i *impl) Name() string {
	return cluster.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	cluster.RegisterRPCServer(server, i)
}
