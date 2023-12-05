package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	cluster "github.com/infraboard/mpaas/apps/k8s"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service cluster.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = logger.Sub(cluster.AppName)
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
