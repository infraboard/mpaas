package api

import (
	"sync"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
)

var (
	h = &handler{}
)

type handler struct {
	service cluster.Service
	log     logger.Logger
	clients map[string]*k8s.Client
	sync.Mutex
}

func (h *handler) Config() error {
	h.clients = map[string]*k8s.Client{}
	h.log = zap.L().Named(cluster.AppName)
	h.service = app.GetGrpcApp(cluster.AppName).(cluster.Service)
	return nil
}

// /prifix/cluster/
func (h *handler) Name() string {
	return proxy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(r *restful.WebService) {
}

func init() {
	app.RegistryRESTfulApp(h)
}
