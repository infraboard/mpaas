package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/deploy"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	deploy.UnimplementedRPCServer

	mcenter *rpc.ClientSet
	cluster cluster.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection("deploys")
	i.log = zap.L().Named(i.Name())
	i.mcenter = rpc.C()
	i.cluster = app.GetInternalApp(cluster.AppName).(cluster.Service)
	return nil
}

func (i *impl) Name() string {
	return deploy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	deploy.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
