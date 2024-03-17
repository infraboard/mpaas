package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mpaas/apps/deploy"
)

func init() {
	ioc.Api().Registry(&downloadHandler{})
}

type downloadHandler struct {
	service deploy.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *downloadHandler) Init() error {
	h.log = log.Sub(deploy.AppName)
	h.service = ioc.Controller().Get(deploy.AppName).(deploy.Service)
	h.Registry()
	return nil
}

func (h *downloadHandler) Name() string {
	return "export"
}

func (h *downloadHandler) Version() string {
	return "v1"
}
