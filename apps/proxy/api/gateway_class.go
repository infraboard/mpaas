package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (h *handler) registryGatewayClassHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] 网关类管理"}

	ws.Route(ws.POST("/{cluster_id}/gateway/classes").To(h.CreateGatewayInstance).
		Doc("创建网关类").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.GatewayClass{}).
		Writes(gatewayv1.GatewayClass{}).
		Returns(200, "OK", gatewayv1.GatewayClass{}))

	ws.Route(ws.GET("/{cluster_id}/gateway/classes").To(h.QueryGatewayInstance).
		Doc("查询网关类列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.ListRequest{}).
		Writes(gatewayv1.GatewayClassList{}).
		Returns(200, "OK", gatewayv1.GatewayClassList{}))

	ws.Route(ws.GET("/{cluster_id}/gateway/classes/{name}").To(h.GetGatewayInstance).
		Doc("查询网关类详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.GetRequest{}).
		Writes(gatewayv1.GatewayClass{}).
		Returns(200, "OK", gatewayv1.GatewayClass{}))

	ws.Route(ws.PUT("/{cluster_id}/gateway/classes/{name}").To(h.UpdateGatewayClass).
		Doc("更新网关类详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.GatewayClass{}).
		Writes(gatewayv1.GatewayClass{}).
		Returns(200, "OK", gatewayv1.GatewayClass{}))

	ws.Route(ws.DELETE("/{cluster_id}/gateway/classes/{name}").To(h.DeleteGatewayClass).
		Doc("更新网关类详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.DeleteRequest{}).
		Writes(gatewayv1.GatewayClass{}).
		Returns(200, "OK", gatewayv1.GatewayClass{}))
}

func (h *handler) CreateGatewayClass(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.GatewayClass{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewCreateRequest()
	ins, err := client.Gateway().CreateGatewayClass(r.Request.Context(), ins, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryGatewayClass(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Gateway().ListGatewayClass(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetGatewayClass(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.Gateway().GetGatewayClass(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateGatewayClass(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.GatewayClass{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Name = r.PathParameter("name")

	ins, err := client.Gateway().UpdateGatewayClass(r.Request.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteGatewayClass(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewDeleteRequest(r.PathParameter("name"))
	ins, err := client.Gateway().DeleteGatewayClass(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
