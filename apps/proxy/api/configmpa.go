package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/binding"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryConfigMapHandler(ws *restful.WebService) {
	tags := []string{"Config Map管理"}
	ws.Route(ws.POST("/{cluster_id}/configmaps").To(h.CreateConfigMap).
		Doc("创建ConfigMap").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(corev1.ConfigMap{}).
		Writes(corev1.ConfigMap{}))

	ws.Route(ws.GET("/{cluster_id}/configmaps").To(h.QueryConfigMap).
		Doc("查询ConfigMap列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.ConfigMapList{}).
		Returns(200, "OK", corev1.ConfigMapList{}))

	ws.Route(ws.GET("/{cluster_id}/configmaps/{name}").To(h.GetConfigMap).
		Doc("查询ConfigMap详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(corev1.ConfigMap{}).
		Returns(200, "OK", corev1.ConfigMap{}))
}

func (h *handler) QueryConfigMap(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Config().ListConfigMap(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetConfigMap(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.Config().GetConfigMap(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) CreateConfigMap(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
	h.log.Debug(tk)

	req := &corev1.ConfigMap{}
	if err := binding.Bind(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充应用关联信息

	// 补充操作人

	ins, err := client.Config().CreateConfigMap(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
