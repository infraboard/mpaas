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
	ioc.RegistryApi(&PipelineTaskHandler{})
}

type PipelineTaskHandler struct {
	service task.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *PipelineTaskHandler) Init() error {
	h.log = zap.L().Named(task.AppName)
	h.service = ioc.GetController(task.AppName).(task.Service)
	return nil
}

func (h *PipelineTaskHandler) Name() string {
	return "pipeline_tasks"
}

func (h *PipelineTaskHandler) Version() string {
	return "v1"
}

func (h *PipelineTaskHandler) APIPrefix() string {
	return fmt.Sprintf("%s/%s/%s",
		conf.C().App.HTTPPrefix(),
		h.Version(),
		h.Name(),
	)
}

func (h *PipelineTaskHandler) Registry(ws *restful.WebService) {
	h.RegistryUserHandler(ws)
}
