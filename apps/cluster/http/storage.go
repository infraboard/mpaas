package http

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s/meta"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryPVHandler(ws *restful.WebService) {
	tags := []string{"存储管理"}

	ws.Route(ws.GET("/{id}/pv").To(h.QueryPersistentVolumes).
		Doc("查询卷列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.PersistentVolumeList{}).
		Returns(200, "OK", corev1.PersistentVolumeList{}))

	ws.Route(ws.GET("/{id}/pvc").To(h.QueryPersistentVolumeClaims).
		Doc("查询PVC列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.PersistentVolumeList{}).
		Returns(200, "OK", corev1.PersistentVolumeList{}))

	ws.Route(ws.GET("/{id}/sc").To(h.QueryStorageClass).
		Doc("查询存储类列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.PersistentVolumeList{}).
		Returns(200, "OK", corev1.PersistentVolumeList{}))
}

func (h *handler) QueryPersistentVolumes(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Storage().ListPersistentVolume(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryPersistentVolumeClaims(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Storage().ListPersistentVolumeClaims(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryStorageClass(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Storage().ListStorageClass(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
