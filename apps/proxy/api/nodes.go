package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryNodeHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] Node管理"}
	ws.Route(ws.GET("/{cluster_id}/nodes").To(h.QueryNodes).
		Doc("查询节点列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.NodeList{}).
		Returns(200, "OK", corev1.NodeList{}))

	ws.Route(ws.GET("/{cluster_id}/nodes/{name}").To(h.GetNode).
		Doc("查询节点详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.Node{}).
		Returns(200, "OK", corev1.Node{}))
}

func (h *handler) QueryNodes(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins, err := client.Admin().ListNode(r.Request.Context(), meta.NewListRequest())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetNode(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.Admin().GetNode(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
