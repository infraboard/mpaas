package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
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

// /prifix/proxy/
func (h *handler) Name() string {
	return proxy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(r *restful.WebService) {
	r.Filter(h.ClusterMiddleware)
	h.registryConfigMapHandler(r)
	h.registryDeploymentHandler(r)
	h.registryNodeHandler(r)
	h.registryNamespaceHandler(r)
	h.registryPodHandler(r)
	h.registrySecretHandler(r)
	h.registryServiceHandler(r)
	h.registryStatefulSetHandler(r)
	h.registryPVHandler(r)
	h.registryWatchHandler(r)
}

// 解析Cluster Id的中间件
func (h *handler) ClusterMiddleware(
	req *restful.Request,
	resp *restful.Response,
	next *restful.FilterChain) {

	// 处理请求
	clusterId := req.PathParameter("cluster_id")
	if clusterId == "" {
		response.Failed(resp, fmt.Errorf("url path param cluster_id required"))
		return
	}

	// 获取集群client对象
	descReq := cluster.NewDescribeClusterRequest(clusterId)
	ins, err := h.service.DescribeCluster(req.Request.Context(), descReq)
	if err != nil {
		response.Failed(resp, fmt.Errorf("describe cluster_id error, %s", err))
		return
	}

	client, err := k8s.NewClient(ins.Spec.KubeConfig)
	if err != nil {
		response.Failed(resp, fmt.Errorf("new k8s client error, %s", err))
		return
	}
	req.SetAttribute(proxy.ATTRIBUTE_K8S_CLIENT, client)

	// next flow
	next.ProcessFilter(req, resp)

	// 处理响应
}
