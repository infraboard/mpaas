package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/task"
)

// 用户自己手动管理任务状态相关接口
func (h *PipelineTaskHandler) RegistryUserHandler(ws *restful.WebService) {
	tags := []string{"Pipeline任务管理"}
	ws.Route(ws.POST("/{id}/status").To(h.QueryPipelineTask).
		Doc("更新任务状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Reads(task.UpdateJobTaskStatusRequest{}).
		Writes(task.JobTask{}))
}

func (h *PipelineTaskHandler) QueryPipelineTask(r *restful.Request, w *restful.Response) {
	req := task.NewQueryPipelineTaskRequest()
	set, err := h.service.QueryPipelineTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
