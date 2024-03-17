package api

import (
	"io"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	"sigs.k8s.io/yaml"

	v1 "k8s.io/api/core/v1"
)

func (h *handler) registrySecretHandler() {
	tags := []string{"[Proxy] 密钥管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/{cluster_id}/secrets").To(h.CreateSecret).
		Doc("创建密钥").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.Secret{}).
		Writes(v1.Secret{}).
		Returns(200, "OK", v1.Secret{}))

	ws.Route(ws.GET("/{cluster_id}/secrets").To(h.QuerySecret).
		Doc("查询密钥列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.Secret{}).
		Writes(v1.SecretList{}).
		Returns(200, "OK", v1.SecretList{}))

	ws.Route(ws.GET("/{cluster_id}/secrets/{name}").To(h.GetSecret).
		Doc("查询密钥详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(v1.Secret{}).
		Writes(v1.Secret{}).
		Returns(200, "OK", v1.Secret{}))
}

func (h *handler) CreateSecret(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

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

	ins, err := client.Config().CreateSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QuerySecret(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Config().ListSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetSecret(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.Config().GetSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
