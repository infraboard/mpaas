package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"github.com/infraboard/mpaas/apps/k8s"
)

func init() {
	ioc.Controller().Registry(&service{})
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
	s.col = ioc_mongo.DB().Collection(s.Name())
	s.encryptoKey = application.Get().EncryptKey
	s.log = log.Sub(s.Name())
	s.cluster = ioc.Controller().Get(k8s.AppName).(k8s.Service)
	return nil
}

func (s *service) Name() string {
	return k8s.AppName
}

func (s *service) Registry(server *grpc.Server) {
	k8s.RegisterRPCServer(server, s)
}
