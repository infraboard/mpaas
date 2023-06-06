package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/task"
)

var (
	h = &handler{}
)

type handler struct {
	service task.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(task.AppName)
	h.service = ioc.GetController(task.AppName).(task.Service)
	return nil
}

func (h *handler) Name() string {
	return task.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	h.RegistryUserHandler(ws)
}

func init() {
	ioc.RegistryApi(h)
}
