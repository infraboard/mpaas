package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/deploy"
)

var (
	h = &handler{}
)

type handler struct {
	service deploy.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(deploy.AppName)
	h.service = ioc.GetController(deploy.AppName).(deploy.Service)
	return nil
}

func (h *handler) Name() string {
	return deploy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	ioc.RegistryApi(h)
}
