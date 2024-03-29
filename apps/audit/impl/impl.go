package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
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
	i.log = log.Sub(i.Name())
	audit.RegisterRPCServer(grpc.Get().Server(), i)
	return nil
}

func (i *impl) Name() string {
	return audit.AppName
}
