package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/proxy"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/meta"

	appsv1 "k8s.io/api/apps/v1"
	scalv1 "k8s.io/api/autoscaling/v1"
)

func (h *handler) registryDeploymentHandler() {
	tags := []string{"[Proxy] Deployment管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/{cluster_id}/deployments").To(h.CreateDeployment).
		Doc("创建Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(appsv1.Deployment{}).
		Writes(appsv1.Deployment{}))

	ws.Route(ws.GET("/{cluster_id}/deployments").To(h.QueryDeployments).
		Doc("查询Deployment列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.Deployment{}).
		Returns(200, "OK", appsv1.Deployment{}))

	ws.Route(ws.GET("/{cluster_id}/deployments/{name}").To(h.GetDeployment).
		Doc("查询Deployment详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.Deployment{}).
		Returns(200, "OK", appsv1.Deployment{}))

	ws.Route(ws.PUT("/{cluster_id}/deployments/{name}").To(h.UpdateDeployment).
		Doc("更新Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.Deployment{}).
		Returns(200, "OK", appsv1.Deployment{}))

	ws.Route(ws.POST("/{cluster_id}/deployments/{name}/scale").To(h.ScaleDeployment).
		Doc("更新副本数").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(scalv1.Scale{}).
		Returns(200, "OK", scalv1.Scale{}))

	ws.Route(ws.POST("/{cluster_id}/deployments/{name}/redeploy").To(h.ReDeployment).
		Doc("重新部署").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(appsv1.Deployment{}).
		Returns(200, "OK", appsv1.Deployment{}))
}

func (h *handler) CreateDeployment(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := &appsv1.Deployment{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.WorkLoad().CreateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryDeployments(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewListRequestFromHttp(r.Request)
	ins, err := client.WorkLoad().ListDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetDeployment(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.WorkLoad().GetDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateDeployment(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := &appsv1.Deployment{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Name = r.PathParameter("name")

	ins, err := client.WorkLoad().UpdateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) ScaleDeployment(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewScaleRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	req.Scale.Name = r.PathParameter("name")

	ins, err := client.WorkLoad().ScaleDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) ReDeployment(r *restful.Request, w *restful.Response) {
	client := r.Attribute(proxy.ATTRIBUTE_K8S_CLIENT).(*k8s.Client)

	req := meta.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")

	ins, err := client.WorkLoad().ReDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
