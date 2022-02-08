package http

import (
	"context"
	"sync"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
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
	service cluster.ServiceServer
	log     logger.Logger
	clients map[string]*k8s.Client
	sync.Mutex
}

func (h *handler) Config() error {
	h.clients = map[string]*k8s.Client{}
	h.log = zap.L().Named(cluster.AppName)
	h.service = app.GetGrpcApp(cluster.AppName).(cluster.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return cluster.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	rr := r.ResourceRouter("cluster")

	// clusters
	rr.BasePath("clusters")
	rr.Handle("POST", "/", h.CreateCluster).AddLabel(label.Create)
	rr.Handle("GET", "/", h.QueryCluster).AddLabel(label.List)
	rr.Handle("GET", "/:id", h.DescribeCluster).AddLabel(label.Get)
	rr.Handle("PUT", "/:id", h.PutCluster).AddLabel(label.Update)
	rr.Handle("PATCH", "/:id", h.PatchCluster).AddLabel(label.Update)
	rr.Handle("DELETE", "/:id", h.DeleteCluster).AddLabel(label.Delete)

	// nodes
	nr := r.ResourceRouter("node")
	nr.BasePath("clusters/:id/nodes")
	nr.Handle("GET", "/", h.QueryNodes).AddLabel(label.List)

	// namespaces
	ns := r.ResourceRouter("namespace")
	ns.BasePath("clusters/:id/namespaces")
	ns.Handle("GET", "/", h.QueryNamespaces).AddLabel(label.List)
	ns.Handle("POST", "/", h.CreateNamespaces).AddLabel(label.Create)

	// deployments
	dr := r.ResourceRouter("deployment")
	dr.BasePath("clusters/:id/deployments")
	dr.Handle("GET", "/", h.QueryDeployments).AddLabel(label.List)
	dr.Handle("POST", "/", h.CreateDeployment).AddLabel(label.Create)
}

func (h *handler) GetClient(ctx context.Context, clusterId string) (*k8s.Client, error) {
	h.Lock()
	defer h.Unlock()

	// 本地缓存中获取
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
	app.RegistryHttpApp(h)
}
