package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"

	cluster "github.com/infraboard/mpaas/apps/k8s"
)

func (h *handler) Registry() {
	tags := []string{"集群管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/").To(h.CreateCluster).
		Doc("创建集群").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.CreateClusterRequest{}).
		Writes(cluster.Cluster{}))

	ws.Route(ws.GET("/").To(h.QueryCluster).
		Doc("查询集群列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(cluster.ClusterSet{}).
		Returns(200, "OK", cluster.ClusterSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeCluster).
		Doc("集群详情").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(cluster.Cluster{}).
		Returns(200, "OK", cluster.Cluster{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutCluster).
		Doc("修改集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(cluster.Cluster{}).
		Returns(200, "OK", cluster.Cluster{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchCluster).
		Doc("修改集群").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(cluster.Cluster{}).
		Returns(200, "OK", cluster.Cluster{}).
		Returns(404, "Not Found", nil))

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
	req.UpdateOwner(token.GetTokenFromRequest(r))

	ins, err := h.service.CreateCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewQueryClusterRequestFromHTTP(r)

	set, err := h.service.QueryCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	set.Desense()
	response.Success(w, set)
}

func (h *handler) DescribeCluster(r *restful.Request, w *restful.Response) {
	req := cluster.NewDescribeClusterRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desense()
	response.Success(w, ins)
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

func (h *handler) PatchCluster(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)
	req := cluster.NewPatchClusterRequest(r.PathParameter("id"))

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
	req := cluster.NewDeleteClusterRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteCluster(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
