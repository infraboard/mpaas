package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mpaas/apps/audit"
	"github.com/infraboard/mpaas/apps/deploy"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service audit.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(deploy.AppName)
	h.service = ioc.Controller().Get(audit.AppName).(audit.Service)
	return nil
}

func (h *handler) Name() string {
	return audit.AppName
}

func (h *handler) Version() string {
	return "v1"
}
