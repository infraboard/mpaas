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

func (h *handler) registryGatewayGrpcRouteHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] 网关GRPC路由管理"}

	ws.Route(ws.POST("/{cluster_id}/{namespace}/gateway/grpcroutes").To(h.CreateGatewayGrpcRoute).
		Doc("创建GRPC路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.GRPCRoute{}).
		Writes(gatewayv1.GRPCRoute{}).
		Returns(200, "OK", gatewayv1.GRPCRoute{}))

	ws.Route(ws.GET("/{cluster_id}/{namespace}/gateway/grpcroutes").To(h.QueryGatewayGrpcRoute).
		Doc("查询GRPC路由列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.ListRequest{}).
		Writes(gatewayv1.GRPCRouteList{}).
		Returns(200, "OK", gatewayv1.GRPCRouteList{}))

	ws.Route(ws.GET("/{cluster_id}/{namespace}/gateway/grpcroutes/{name}").To(h.GetGatewayGrpcRoute).
		Doc("查询GRPC路由详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.GetRequest{}).
		Writes(gatewayv1.GRPCRoute{}).
		Returns(200, "OK", gatewayv1.GRPCRoute{}))

	ws.Route(ws.PUT("/{cluster_id}/{namespace}/gateway/httproutes/{name}").To(h.UpdateGrpcRoute).
		Doc("更新GRPC路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.GRPCRoute{}).
		Writes(gatewayv1.GRPCRoute{}).
		Returns(200, "OK", gatewayv1.GRPCRoute{}))

	ws.Route(ws.DELETE("/{cluster_id}/{namespace}/gateway/httproutes/{name}").To(h.DeleteGrpcRoute).
		Doc("删除GRPC路由").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(gatewayv1.GRPCRoute{}).
		Returns(200, "OK", gatewayv1.GRPCRoute{}))
}

func (h *handler) CreateGatewayGrpcRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.GRPCRoute{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Namespace = r.PathParameter("namespace")

	req := meta.NewCreateRequest()
	ins, err := client.Gateway().CreateGRPCRoute(r.Request.Context(), ins, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryGatewayGrpcRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.Gateway().ListGRPCRouteList(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetGatewayGrpcRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.Gateway().GetGRPCRoute(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateGrpcRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.GRPCRoute{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Namespace = r.PathParameter("namespace")
	ins.Name = r.PathParameter("name")

	ins, err := client.Gateway().UpdateGRPCRoute(r.Request.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteGrpcRoute(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewDeleteRequest(r.PathParameter("name"))
	req.Namespace = r.PathParameter("namespace")
	ins, err := client.Gateway().DeleteGRPCRoute(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
