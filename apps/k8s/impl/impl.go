package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col     *mongo.Collection
	log     logger.Logger
	cluster cluster.Service
	cluster.UnimplementedRPCServer
	ioc.IocObjectImpl
	encryptoKey string
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(s.Name())

	s.encryptoKey = conf.C().App.EncryptKey
	s.log = zap.L().Named(s.Name())
	s.cluster = ioc.GetController(cluster.AppName).(cluster.Service)
	return nil
}

func (s *service) Name() string {
	return cluster.AppName
}

func (s *service) Registry(server *grpc.Server) {
	cluster.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
