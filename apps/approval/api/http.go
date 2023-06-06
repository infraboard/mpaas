package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mpaas/apps/approval"
)

var (
	h = &handler{}
)

type handler struct {
	service approval.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(approval.AppName)
	h.service = ioc.GetController(approval.AppName).(approval.Service)
	return nil
}

// /prifix/cluster/
func (h *handler) Name() string {
	return approval.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func init() {
	ioc.RegistryApi(h)
}
