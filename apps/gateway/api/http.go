package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
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
	h.log = log.Sub(gateway.AppName)
	h.service = ioc.Controller().Get(gateway.AppName).(gateway.Service)
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
	ioc.Api().Registry(h)
}
