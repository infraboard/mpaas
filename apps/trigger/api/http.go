package api

import (
	"fmt"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/conf"
)

var (
	h = &Handler{}
)

type Handler struct {
	log logger.Logger
	svc trigger.Service
}

func (h *Handler) Config() error {
	h.svc = app.GetInternalApp(trigger.AppName).(trigger.Service)
	h.log = zap.L().Named(trigger.AppName)
	return nil
}

func (h *Handler) Name() string {
	return trigger.AppName
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
	tags := []string{"事件处理"}

	ws.Route(ws.POST("gitlab").To(h.HandleGitlabEvent).
		Doc("处理Gitlab Webhook事件").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRESTfulApp(h)
}
