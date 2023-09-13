package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务集群管理"}
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
}

func (h *handler) CreateCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewCreateClusterRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateCluster(r.Request.Context(), req)
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
	response.Success(w, set)
}
