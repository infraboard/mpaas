package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
)

func (h *handler) Registry() {
	tags := []string{"服务集群管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/").To(h.CreateCluster).
		Doc("创建集群").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.CreateDeploymentRequest{}).
		Writes(deploy.Deployment{}))

	ws.Route(ws.GET("/").To(h.QueryCluster).
		Doc("查询集群列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.QueryDeploymentRequest{}).
		Writes(deploy.DeploymentSet{}).
		Returns(200, "OK", deploy.DeploymentSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeCluster).
		Doc("查询集群详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.DescribeDeploymentRequest{}).
		Writes(deploy.Deployment{}).
		Returns(200, "OK", deploy.Deployment{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteCluster).
		Doc("删除集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))
}

func (h *handler) CreateCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewCreateClusterRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateCluster(token.WithTokenCtx(r), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewQueryClusterRequestFromHttp(r)

	set, err := h.service.QueryCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 针对前端专门做Tree转换
	if r.QueryParameter("to_tree") == "true" {
		response.Success(w, ClusterSetToTreeSet(set))
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewDescribeClusterRequest(r.PathParameter("id"))

	set, err := h.service.DescribeCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) PutCluster(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)

	req := cluster.NewPutClusterRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewDeleteClusterRequest(r.PathParameter("id"))
	set, err := h.service.DeleteCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
