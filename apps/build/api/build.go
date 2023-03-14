package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/build"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"构建配置管理"}
	ws.Route(ws.POST("/").To(h.CreateBuildConfig).
		Doc("创建构建配置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(build.CreateBuildConfigRequest{}).
		Writes(build.BuildConfig{}))

	ws.Route(ws.GET("/").To(h.QueryBuildConfig).
		Doc("查询构建配置列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Disable).
		Metadata(label.Permission, label.Enable).
		Reads(build.QueryBuildConfigRequest{}).
		Writes(build.BuildConfigSet{}).
		Returns(200, "OK", build.BuildConfigSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeBuildConfig).
		Doc("构建配置详情").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Disable).
		Metadata(label.Permission, label.Enable).
		Writes(build.BuildConfig{}).
		Returns(200, "OK", build.BuildConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutBuildConfig).
		Doc("修改构建配置").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(build.BuildConfig{}).
		Returns(200, "OK", build.BuildConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PatchBuildConfig).
		Doc("修改构建配置").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(build.BuildConfig{}).
		Returns(200, "OK", build.BuildConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteBuildConfig).
		Doc("删除构建配置").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))

}

func (h *handler) CreateBuildConfig(r *restful.Request, w *restful.Response) {
	req := build.NewCreateBuildConfigRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryBuildConfig(r *restful.Request, w *restful.Response) {
	req := build.NewQueryBuildConfigRequestFromHTTP(r.Request)
	set, err := h.service.QueryBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeBuildConfig(r *restful.Request, w *restful.Response) {
	req := build.NewDescribeBuildConfigRequst(r.PathParameter("id"))
	ins, err := h.service.DescribeBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PutBuildConfig(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)

	req := build.NewPutBuildConfigRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchBuildConfig(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)

	req := build.NewPatchBuildConfigRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteBuildConfig(r *restful.Request, w *restful.Response) {
	req := build.NewDeleteBuildConfigRequest(r.PathParameter("id"))
	set, err := h.service.DeleteBuildConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
