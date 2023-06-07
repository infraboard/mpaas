package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/build"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service build.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(build.AppName)
	h.service = ioc.GetController(build.AppName).(build.Service)
	return nil
}

func (h *handler) Name() string {
	return build.AppName
}

func (h *handler) Version() string {
	return "v1"
}
