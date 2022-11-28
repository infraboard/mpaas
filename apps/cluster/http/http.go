package http

import (
	"context"
	"sync"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/cluster"
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
	return cluster.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(r *restful.WebService) {
	h.registryClusterHandler(r)
	h.registryNodeHandler(r)
	h.registryNamespaceHandler(r)
	h.registryDeploymentHandler(r)
	h.registryPodHandler(r)
	h.registryWatchHandler(r)
	h.registryConfigMapHandler(r)
	h.registryPVHandler(r)
	h.registryServiceHandler(r)
	h.registrySecretHandler(r)
	h.registryStatefulSetHandler(r)
}

func (h *handler) GetClient(ctx context.Context, clusterId string) (*k8s.Client, error) {
	h.Lock()
	defer h.Unlock()

	// 本地缓存中获取, 当Client有更新时，需要更新缓存
	v, ok := h.clients[clusterId]
	if ok {
		return v, nil
	}

	req := cluster.NewDescribeClusterRequest(clusterId)
	ins, err := h.service.DescribeCluster(ctx, req)
	if err != nil {
		return nil, err
	}

	client, err := k8s.NewClient(ins.Data.KubeConfig)
	if err != nil {
		return nil, err
	}

	h.clients[ins.Id] = client
	return client, nil
}

func init() {
	app.RegistryRESTfulApp(h)
}
