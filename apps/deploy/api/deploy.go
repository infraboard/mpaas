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
	ws.Route(ws.POST("/").To(h.CreateDeployConfig).
		Doc("部署配置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.CreateDeployConfigRequest{}).
		Writes(deploy.DeployConfig{}))

	ws.Route(ws.GET("/").To(h.QueryDeployConfig).
		Doc("查询部署配置列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(deploy.QueryDeployConfigRequest{}).
		Writes(deploy.DeployConfigSet{}).
		Returns(200, "OK", deploy.DeployConfigSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeDeployConfig).
		Doc("部署配置详情").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.DeployConfig{}).
		Returns(200, "OK", deploy.DeployConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDeployConfig).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.DeployConfig{}).
		Returns(200, "OK", deploy.DeployConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDeployConfig).
		Doc("修改部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(deploy.DeployConfig{}).
		Returns(200, "OK", deploy.DeployConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteDeployConfig).
		Doc("删除部署配置").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable))
}

func (h *handler) CreateDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewCreateDeployConfigRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewQueryDeployConfigRequestFromHttp(r.Request)

	set, err := h.service.QueryDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeployConfigRequest(r.PathParameter("id"))

	ins, err := h.service.DescribeDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) DownloadDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeployConfigRequest(r.PathParameter("id"))

	// 查询部署配置
	ins, err := h.service.DescribeDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 校验集群访问Token
	if ins.Spec.AuthEnabled {
		err = ins.ValidateToken(r.QueryParameter(deploy.DEPLOY_CONFIG_ACCESS_TOKEN_HEADER))
		if err != nil {
			response.Failed(w, err)
			return
		}
	}

	switch ins.Spec.Type {
	case deploy.TYPE_HOST:
	case deploy.TYPE_KUBERNETES:
		w.Write([]byte(ins.Spec.K8STypeConfig.WorkloadConfig))
	}
}

func (h *handler) PutDeployConfig(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPutDeployRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchDeployConfig(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := deploy.NewPatchDeployRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdateDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewDeleteDeployConfigRequest(r.PathParameter("id"))
	set, err := h.service.DeleteDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
