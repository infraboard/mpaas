package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/conf"
)

func init() {
	ioc.RegistryApi(&Handler{})
}

type Handler struct {
	service task.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *Handler) Init() error {
	h.log = zap.L().Named(task.AppName)
	h.service = ioc.GetController(task.AppName).(task.Service)
	return nil
}

func (h *Handler) Name() string {
	return task.AppName
}

func (h *Handler) Version() string {
	return "v1"
}

func (h *Handler) APIPrefix() string {
	return fmt.Sprintf("%s/%s/%s",
		conf.C().App.HTTPPrefix(),
		h.Version(),
		h.Name(),
	)
}

func (h *Handler) Registry(ws *restful.WebService) {
	h.RegistryUserHandler(ws)
}
