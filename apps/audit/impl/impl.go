package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
	"github.com/infraboard/mpaas/apps/audit"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	audit.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())
	return nil
}

func (i *impl) Name() string {
	return audit.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	audit.RegisterRPCServer(server, i)
}
