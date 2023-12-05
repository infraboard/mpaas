package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/clients/rpc"
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
	deploy.UnimplementedRPCServer
	ioc.ObjectImpl

	mcenter *rpc.ClientSet
	k8s     k8s.Service
	cluster cluster.Service
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())
	i.mcenter = rpc.C()
	i.k8s = ioc.GetController(k8s.AppName).(k8s.Service)
	i.cluster = ioc.GetController(cluster.AppName).(cluster.Service)
	return nil
}

func (i *impl) Name() string {
	return deploy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	deploy.RegisterRPCServer(server, i)
}
