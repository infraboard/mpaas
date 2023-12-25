package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	"github.com/infraboard/mpaas/apps/k8s"
)

func init() {
	ioc.Controller().Registry(&impl{})
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
	i.log = log.Sub(i.Name())
	i.mcenter = rpc.C()
	i.k8s = ioc.Controller().Get(k8s.AppName).(k8s.Service)
	i.cluster = ioc.Controller().Get(cluster.AppName).(cluster.Service)
	return nil
}

func (i *impl) Name() string {
	return deploy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	deploy.RegisterRPCServer(server, i)
}
