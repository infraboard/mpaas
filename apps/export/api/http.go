package api

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/deploy"
)

var (
	dh = &downloadHandler{}
)

type downloadHandler struct {
	service deploy.Service
	log     logger.Logger
}

func (h *downloadHandler) Config() error {
	h.log = zap.L().Named(deploy.AppName)
	h.service = app.GetGrpcApp(deploy.AppName).(deploy.Service)
	return nil
}

func (h *downloadHandler) Name() string {
	return "export"
}

func (h *downloadHandler) Version() string {
	return "v1"
}

func init() {
	app.RegistryRESTfulApp(dh)
}
