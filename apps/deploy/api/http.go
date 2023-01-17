package api

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/deploy"
)

var (
	h = &handler{}
)

type handler struct {
	service deploy.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(deploy.AppName)
	h.service = app.GetGrpcApp(deploy.AppName).(deploy.Service)
	return nil
}

// /prifix/cluster/
func (h *handler) Name() string {
	return deploy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	app.RegistryRESTfulApp(h)
}
