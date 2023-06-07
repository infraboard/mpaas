package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/job"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service job.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(job.AppName)
	h.service = ioc.GetController(job.AppName).(job.Service)
	return nil
}

func (h *handler) Name() string {
	return job.AppName
}

func (h *handler) Version() string {
	return "v1"
}
