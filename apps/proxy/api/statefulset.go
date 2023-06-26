package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"

	appsv1 "k8s.io/api/apps/v1"
)

func (h *handler) registryStatefulSetHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] StatefulSet管理"}
	ws.Route(ws.POST("/{cluster_id}/statefulsets").To(h.CreateStatefulSet).
		Doc("创建StatefulSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(appsv1.StatefulSet{}).
		Writes(appsv1.StatefulSet{}))

	ws.Route(ws.GET("/{cluster_id}/statefulsets").To(h.QueryStatefulSet).
		Doc("查询StatefulSet列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.StatefulSetList{}).
		Returns(200, "OK", appsv1.StatefulSetList{}))

	ws.Route(ws.GET("/{cluster_id}/statefulsets/{name}").To(h.GetStatefulSet).
		Doc("查询StatefulSet详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.StatefulSet{}).
		Returns(200, "OK", appsv1.StatefulSet{}))
}

func (h *handler) CreateStatefulSet(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := &appsv1.StatefulSet{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.WorkLoad().CreateStatefulSet(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryStatefulSet(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.WorkLoad().ListStatefulSet(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetStatefulSet(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.WorkLoad().GetStatefulSet(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
