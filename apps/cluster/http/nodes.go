package http

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"
	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryNodeHandler(ws *restful.WebService) {
	tags := []string{"Node管理"}
	ws.Route(ws.GET("/").To(h.QueryNodes).
		Doc("查询Namespace").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(corev1.Node{})).
		Returns(200, "OK", corev1.Node{}))
}

func (h *handler) QueryNodes(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.ListNode(r.Request.Context(), k8s.NewListRequest())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
