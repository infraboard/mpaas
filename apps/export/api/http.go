package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/deploy"
)

var (
	dh = &downloadHandler{}
)

type downloadHandler struct {
	service deploy.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *downloadHandler) Init() error {
	h.log = zap.L().Named(deploy.AppName)
	h.service = ioc.GetController(deploy.AppName).(deploy.Service)
	return nil
}

func (h *downloadHandler) Name() string {
	return "export"
}

func (h *downloadHandler) Version() string {
	return "v1"
}

func init() {
	ioc.RegistryApi(dh)
}
