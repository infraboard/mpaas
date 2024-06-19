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

func (h *handler) registryGatewayHttpRouteHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] 网关HTTP路由管理"}

	ws.Route(ws.POST("/{cluster_id}/gateway/httproutes").To(h.CreateGatewayHttpRoute).
		Doc("创建HTTP路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.HTTPRoute{}).
		Writes(gatewayv1.HTTPRoute{}).
		Returns(200, "OK", gatewayv1.HTTPRoute{}))

	ws.Route(ws.GET("/{cluster_id}/gateway/httproutes/{namespace}").To(h.QueryGatewayHttpRoute).
		Doc("查询HTTP路由列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.ListRequest{}).
		Writes(gatewayv1.HTTPRouteList{}).
		Returns(200, "OK", gatewayv1.HTTPRouteList{}))

	ws.Route(ws.GET("/{cluster_id}/gateway/httproutes/{namespace}/{name}").To(h.GetGatewayHttpRoute).
		Doc("查询HTTP路由详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.GetRequest{}).
		Writes(gatewayv1.HTTPRoute{}).
		Returns(200, "OK", gatewayv1.HTTPRoute{}))

	ws.Route(ws.PUT("/{cluster_id}/gateway/httproutes/{namespace}/{name}").To(h.UpdateHttpRoute).
		Doc("更新HTTP路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.HTTPRoute{}).
		Writes(gatewayv1.HTTPRoute{}).
		Returns(200, "OK", gatewayv1.HTTPRoute{}))

	ws.Route(ws.DELETE("/{cluster_id}/gateway/httproutes/{namespace}/{name}").To(h.DeleteHttpRoute).
		Doc("删除HTTP路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(gatewayv1.HTTPRoute{}).
		Returns(200, "OK", gatewayv1.HTTPRoute{}))
}

func (h *handler) CreateGatewayHttpRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.HTTPRoute{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewCreateRequest()
	ins, err := client.Gateway().CreateHttpRoute(r.Request.Context(), ins, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryGatewayHttpRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	req.Namespace = r.PathParameter("namespace")
	ins, err := client.Gateway().ListHttpRoute(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetGatewayHttpRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Namespace = r.PathParameter("namespace")
	req.Name = r.PathParameter("name")
	ins, err := client.Gateway().GetHttpRoute(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateHttpRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.HTTPRoute{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Namespace = r.PathParameter("namespace")
	ins.Name = r.PathParameter("name")

	ins, err := client.Gateway().UpdateHttpRoute(r.Request.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteHttpRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewDeleteRequest(r.PathParameter("name"))
	req.Namespace = r.PathParameter("namespace")
	ins, err := client.Gateway().DeleteHttpRoute(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
