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
	ws.Route(ws.POST("/").To(h.CreateDeploy).
		Doc("部署配置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.CreateDeployRequest{}).
		Writes(deploy.Deploy{}))

	ws.Route(ws.GET("/").To(h.QueryDeploy).
		Doc("查询部署配置列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.QueryDeployRequest{}).
		Writes(deploy.DeploySet{}).
		Returns(200, "OK", deploy.DeploySet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeDeploy).
		Doc("部署配置详情").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deploy{}).
		Returns(200, "OK", deploy.Deploy{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDeploy).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deploy{}).
		Returns(200, "OK", deploy.Deploy{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDeploy).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.Deploy{}).
		Returns(200, "OK", deploy.Deploy{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteDeploy).
		Doc("删除部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))
}

func (h *handler) CreateDeploy(r *restful.Request, w *restful.Response) {
	req := deploy.NewCreateDeployRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryDeploy(r *restful.Request, w *restful.Response) {
	req := deploy.NewQueryDeployRequestFromHttp(r.Request)

	set, err := h.service.QueryDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeDeploy(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeployRequest(r.PathParameter("id"))

	ins, err := h.service.DescribeDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PutDeploy(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPutDeployRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchDeploy(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPatchDeployRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteDeploy(r *restful.Request, w *restful.Response) {
	req := deploy.NewDeleteDeployRequest(r.PathParameter("id"))
	set, err := h.service.DeleteDeploy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
