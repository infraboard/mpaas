package api

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/pipeline"
)

var (
	h = &handler{}
)

type handler struct {
	service pipeline.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(pipeline.AppName)
	h.service = app.GetGrpcApp(pipeline.AppName).(pipeline.Service)
	return nil
}

func (h *handler) Name() string {
	return pipeline.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	app.RegistryRESTfulApp(h)
}
