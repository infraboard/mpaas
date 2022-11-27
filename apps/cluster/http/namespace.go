package http

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

func (h *handler) registryNamespaceHandler(ws *restful.WebService) {
	tags := []string{"Namespace管理"}
	ws.Route(ws.POST("/{id}/namespace").To(h.CreateNamespaces).
		Doc("创建Namespace").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(corev1.Namespace{}).
		Writes(response.NewData(corev1.Namespace{})))

	ws.Route(ws.GET("/{id}/namespaces").To(h.QueryNamespaces).
		Doc("查询Namespace").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(corev1.Namespace{})).
		Returns(200, "OK", corev1.Namespace{}))
}

func (h *handler) QueryNamespaces(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := client.ListNamespace(r.Request.Context(), k8s.NewListRequest())
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateNamespaces(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &v1.Namespace{}
	if err := binding.Bind(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
