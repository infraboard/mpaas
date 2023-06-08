package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	job.UnimplementedRPCServer
	ioc.IocObjectImpl
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
	return job.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	job.RegisterRPCServer(server, i)
}
