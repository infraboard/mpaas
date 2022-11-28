package http

import (
	"io"

	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"
	"sigs.k8s.io/yaml"

	v1 "k8s.io/api/core/v1"
)

func (h *handler) registrySecretHandler(ws *restful.WebService) {
	tags := []string{"密钥管理"}

	ws.Route(ws.POST("/{id}/secrets").To(h.CreateService).
		Doc("创建密钥").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.Secret{})).
		Returns(200, "OK", v1.Secret{}))

	ws.Route(ws.GET("/{id}/secrets").To(h.QueryService).
		Doc("查询密钥列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.SecretList{})).
		Returns(200, "OK", v1.SecretList{}))

	ws.Route(ws.GET("/{id}/secrets/{name}").To(h.GetService).
		Doc("查询密钥详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.Secret{})).
		Returns(200, "OK", v1.Secret{}))
}

func (h *handler) CreateSecret(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &v1.Secret{}

	data, err := io.ReadAll(r.Request.Body)
	if err != nil {
		response.Failed(w, err)
		return
	}
	defer r.Request.Body.Close()

	if err := yaml.Unmarshal(data, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QuerySecret(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListRequestFromHttp(r.Request)
	ins, err := client.ListSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetSecret(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.GetService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
