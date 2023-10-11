package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/deploy"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service deploy.Service
	log     logger.Logger
	ioc.ObjectImpl
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
