package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/k8s"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	cluster.UnimplementedRPCServer
	ioc.ObjectImpl

	k8s    k8s.Service
	deploy deploy.RPCServer
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())
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
