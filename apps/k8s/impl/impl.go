package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.RegistryController(&service{})
}

type service struct {
	col     *mongo.Collection
	log     *zerolog.Logger
	cluster k8s.Service
	k8s.UnimplementedRPCServer
	ioc.ObjectImpl
	encryptoKey string
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(s.Name())

	s.encryptoKey = conf.C().App.EncryptKey
	s.log = logger.Sub(s.Name())
	s.cluster = ioc.GetController(k8s.AppName).(k8s.Service)
	return nil
}

func (s *service) Name() string {
	return k8s.AppName
}

func (s *service) Registry(server *grpc.Server) {
	k8s.RegisterRPCServer(server, s)
}
