package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	cluster "github.com/infraboard/mpaas/apps/k8s"
)

var (
	h = &handler{}
)

type handler struct {
	service cluster.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(cluster.AppName)
	h.service = ioc.GetController(cluster.AppName).(cluster.Service)
	return nil
}

// /prifix/cluster/
func (h *handler) Name() string {
	return cluster.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	ioc.RegistryApi(h)
}
