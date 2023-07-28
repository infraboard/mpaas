package api

import (
	"fmt"

	"github.com/infraboard/mcenter/apps/token"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/pipeline"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"Pipeline管理"}
	fmt.Println(tags)
	ws.Route(ws.POST("/").To(h.CreatePipeline).
		Doc("创建Pipeline").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(pipeline.CreatePipelineRequest{}).
		Writes(pipeline.Pipeline{}))

	ws.Route(ws.GET("/").To(h.QueryPipeline).
		Doc("查询Pipeline列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(pipeline.QueryPipelineRequest{}).
		Writes(pipeline.PipelineSet{}).
		Returns(200, "OK", pipeline.PipelineSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribePipeline).
		Doc("Pipeline详情").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(pipeline.Pipeline{}).
		Returns(200, "OK", pipeline.Pipeline{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutPipeline).
		Doc("修改Pipeline").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(pipeline.Pipeline{}).
		Returns(200, "OK", pipeline.Pipeline{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}").To(h.PatchPipeline).
		Doc("修改Pipeline").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(pipeline.Pipeline{}).
		Returns(200, "OK", pipeline.Pipeline{}).
		Returns(404, "Not Found", nil))
}

func (h *handler) CreatePipeline(r *restful.Request, w *restful.Response) {
	req := pipeline.NewCreatePipelineRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UpdateFromToken(token.GetTokenFromRequest(r))
	ins, err := h.service.CreatePipeline(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryPipeline(r *restful.Request, w *restful.Response) {
	req := pipeline.NewQueryPipelineRequestFromHTTP(r)

	set, err := h.service.QueryPipeline(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribePipeline(r *restful.Request, w *restful.Response) {
	req := pipeline.NewDescribePipelineRequest(r.PathParameter("id"))
	ins, err := h.service.DescribePipeline(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PutPipeline(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)

	req := pipeline.NewPutPipelineRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdatePipeline(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchPipeline(r *restful.Request, w *restful.Response) {
	tk := r.Attribute("token").(*token.Token)

	req := pipeline.NewPatchPipelineRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateBy = tk.Username

	set, err := h.service.UpdatePipeline(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
