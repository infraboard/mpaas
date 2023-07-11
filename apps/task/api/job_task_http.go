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
	ioc.RegistryApi(&JobTaskHandler{})
}

type JobTaskHandler struct {
	service task.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *JobTaskHandler) Init() error {
	h.log = zap.L().Named(task.AppName)
	h.service = ioc.GetController(task.AppName).(task.Service)
	return nil
}

func (h *JobTaskHandler) Name() string {
	return "job_tasks"
}

func (h *JobTaskHandler) Version() string {
	return "v1"
}

func (h *JobTaskHandler) APIPrefix() string {
	return fmt.Sprintf("%s/%s/%s",
		conf.C().App.HTTPPrefix(),
		h.Version(),
		h.Name(),
	)
}

func (h *JobTaskHandler) Registry(ws *restful.WebService) {
	h.RegistryUserHandler(ws)
}
