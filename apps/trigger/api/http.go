package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/trigger"
)

var (
	h = &handler{}
)

type handler struct {
	log logger.Logger
	scm trigger.Service
}

func (h *handler) Config() error {
	h.scm = app.GetInternalApp(trigger.AppName).(trigger.Service)
	h.log = zap.L().Named(trigger.AppName)
	return nil
}

func (h *handler) Name() string {
	return trigger.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"事件处理"}

	ws.Route(ws.GET("gitlab").To(h.HandleGitlabEvent).
		Doc("处理Gitlab Webhook事件").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRESTfulApp(h)
}
