package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/gateway"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	gateway.UnimplementedRPCServer
	ioc.IocObjectImpl
}

func (s *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc

	return nil
}

func (s *impl) Name() string {
	return gateway.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	gateway.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
