package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col     *mongo.Collection
	log     logger.Logger
	cluster cluster.ServiceServer
	cluster.UnimplementedServiceServer
}

func (s *service) Config() error {

	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	s.cluster = app.GetGrpcApp(cluster.AppName).(cluster.ServiceServer)
	return nil
}

func (s *service) Name() string {
	return cluster.AppName
}

func (s *service) Registry(server *grpc.Server) {
	cluster.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
