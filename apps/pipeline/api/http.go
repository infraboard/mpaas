package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/pipeline"
)

var (
	h = &handler{}
)

type handler struct {
	service pipeline.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(pipeline.AppName)
	h.service = ioc.GetController(pipeline.AppName).(pipeline.Service)
	return nil
}

func (h *handler) Name() string {
	return pipeline.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	ioc.RegistryApi(h)
}
