package api

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/build"
)

var (
	h = &handler{}
)

type handler struct {
	service build.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(build.AppName)
	h.service = app.GetGrpcApp(build.AppName).(build.Service)
	return nil
}

func (h *handler) Name() string {
	return build.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	app.RegistryRESTfulApp(h)
}
