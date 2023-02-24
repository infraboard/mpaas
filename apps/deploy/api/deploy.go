package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/deploy"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"部署配置管理"}
	ws.Route(ws.POST("/").To(h.CreateDeployment).
		Doc("部署配置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.CreateDeploymentRequest{}).
		Writes(deploy.Deployment{}))

	ws.Route(ws.GET("/").To(h.QueryDeployment).
		Doc("查询部署配置列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.QueryDeploymentRequest{}).
		Writes(deploy.DeploymentSet{}).
		Returns(200, "OK", deploy.DeploymentSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeDeployment).
		Doc("部署配置详情").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deployment{}).
		Returns(200, "OK", deploy.Deployment{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDeployment).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deployment{}).
		Returns(200, "OK", deploy.Deployment{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDeployment).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deployment{}).
		Returns(200, "OK", deploy.Deployment{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteDeployment).
		Doc("删除部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))
}

func (h *handler) CreateDeployment(r *restful.Request, w *restful.Response) {
	req := deploy.NewCreateDeploymentRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryDeployment(r *restful.Request, w *restful.Response) {
	req := deploy.NewQueryDeploymentRequestFromHttp(r.Request)

	set, err := h.service.QueryDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeDeployment(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeploymentRequest(r.PathParameter("id"))

	ins, err := h.service.DescribeDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PutDeployment(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPutDeployRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchDeployment(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPatchDeployRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteDeployment(r *restful.Request, w *restful.Response) {
	req := deploy.NewDeleteDeploymentRequest(r.PathParameter("id"))
	set, err := h.service.DeleteDeployment(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
