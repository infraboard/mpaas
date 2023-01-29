package api

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/job"
)

var (
	h = &handler{}
)

type handler struct {
	service job.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(job.AppName)
	h.service = app.GetGrpcApp(job.AppName).(job.Service)
	return nil
}

func (h *handler) Name() string {
	return job.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	app.RegistryRESTfulApp(h)
}
