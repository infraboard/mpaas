package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mpaas/apps/deploy"
)

func init() {
	ioc.RegistryApi(&downloadHandler{})
}

type downloadHandler struct {
	service deploy.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *downloadHandler) Init() error {
	h.log = logger.Sub(deploy.AppName)
	h.service = ioc.GetController(deploy.AppName).(deploy.Service)
	return nil
}

func (h *downloadHandler) Name() string {
	return "export"
}

func (h *downloadHandler) Version() string {
	return "v1"
}
