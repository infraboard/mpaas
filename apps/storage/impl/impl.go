package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mpaas/apps/storage"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log logger.Logger
	db  *mongo.Database
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.db = db
	s.log = zap.L().Named("storage")
	return nil
}

func (s *service) Name() string {
	return storage.AppName
}

func init() {
	app.RegistryInternalApp(svr)
}
