package http

import (
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/cluster"
)

var (
	h = &handler{}
)

type handler struct {
	service cluster.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(cluster.AppName)
	h.service = app.GetGrpcApp(cluster.AppName).(cluster.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return cluster.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	rr := r.ResourceRouter("cluster")

	rr.BasePath("clusters")
	rr.Handle("POST", "/", h.CreateCluster).AddLabel(label.Create)
	rr.Handle("GET", "/", h.QueryCluster).AddLabel(label.List)
	rr.Handle("GET", "/:id", h.DeleteCluster).AddLabel(label.Get)
	rr.Handle("PUT", "/:id", h.PutCluster).AddLabel(label.Update)
	rr.Handle("PATCH", "/:id", h.PatchCluster).AddLabel(label.Update)
	rr.Handle("DELETE", "/:id", h.DeleteCluster).AddLabel(label.Delete)
}

func init() {
	app.RegistryHttpApp(h)
}
