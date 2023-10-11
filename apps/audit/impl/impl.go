package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/audit"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	audit.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return audit.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	audit.RegisterRPCServer(server, i)
}
