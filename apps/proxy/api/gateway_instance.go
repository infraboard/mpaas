package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (h *handler) registryGatewayInstanceHandler(ws *restful.WebService) {
	tags := []string{"[Proxy] 网关实例管理"}

	ws.Route(ws.POST("/{cluster_id}/{namespace}/gateway/instances").To(h.CreateGatewayInstance).
		Doc("创建网关").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(gatewayv1.Gateway{}).
		Returns(200, "OK", gatewayv1.Gateway{}))

	ws.Route(ws.GET("/{cluster_id}/{namespace}/gateway/instances").To(h.QueryGatewayInstance).
		Doc("查询网关列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(gatewayv1.GatewayList{}).
		Returns(200, "OK", gatewayv1.GatewayList{}))

	ws.Route(ws.GET("/{cluster_id}/{namespace}/gateway/instances/{name}").To(h.GetGatewayInstance).
		Doc("查询网关详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(gatewayv1.Gateway{}).
		Returns(200, "OK", gatewayv1.Gateway{}))

	ws.Route(ws.PUT("/{cluster_id}/{namespace}/gateway/instances/{name}").To(h.UpdateGatewayInstance).
		Doc("更新网关").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(gatewayv1.Gateway{}).
		Writes(gatewayv1.Gateway{}).
		Returns(200, "OK", gatewayv1.Gateway{}))

	ws.Route(ws.DELETE("/{cluster_id}/{namespace}/gateway/instances/{name}").To(h.DeleteGatewayInstance).
		Doc("删除网关").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(meta.DeleteRequest{}).
		Writes(gatewayv1.Gateway{}).
		Returns(200, "OK", gatewayv1.Gateway{}))
}

func (h *handler) CreateGatewayInstance(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.Gateway{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}

	req := meta.NewCreateRequest()
	ins, err := client.Gateway().CreateGateway(r.Request.Context(), ins, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryGatewayInstance(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	req.Namespace = r.PathParameter("namespace")
	ins, err := client.Gateway().ListGateway(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetGatewayInstance(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Namespace = r.PathParameter("namespace")
	req.Name = r.PathParameter("name")
	ins, err := client.Gateway().GetGateway(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateGatewayInstance(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	ins := &gatewayv1.Gateway{}
	if err := r.ReadEntity(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Name = r.PathParameter("name")
	ins.Namespace = r.PathParameter("namespace")

	ins, err := client.Gateway().UpdateGateway(r.Request.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteGatewayInstance(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewDeleteRequest(r.PathParameter("name"))
	req.Namespace = r.PathParameter("namespace")
	ins, err := client.Gateway().DeleteGateway(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
