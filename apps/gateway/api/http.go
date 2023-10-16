package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mpaas/apps/gateway"
)

var (
	h = &handler{}
)

type handler struct {
	service gateway.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = logger.Sub(gateway.AppName)
	h.service = ioc.GetController(gateway.AppName).(gateway.Service)
	return nil
}

func (h *handler) Name() string {
	return gateway.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"网关管理"}
	fmt.Println(tags)
}

func init() {
	ioc.RegistryApi(h)
}
