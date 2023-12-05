package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/mcube/v2/ioc"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"github.com/infraboard/mpaas/apps/gateway"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	gateway.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (s *impl) Init() error {
	dc := ioc_mongo.DB().Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "create_at", Value: -1}},
		},
	}

	_, err := dc.Indexes().CreateMany(context.Background(), indexs)
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
	gateway.RegisterRPCServer(server, s)
}
