package api

import (
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/clients/rpc/middleware"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/common/terminal"
)

var (
	JOB_TASK_RESOURCE_NAME = "JobTask"
)

// 用户自己手动管理任务状态相关接口
func (h *JobTaskHandler) RegistryUserHandler(ws *restful.WebService) {
	tags := []string{"Job任务管理"}
	ws.Route(ws.GET("/").
		To(h.QueryJobTask).
		Doc("查询JobTask运行任务列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(task.QueryJobTaskRequest{}).
		Writes(task.JobTaskSet{}))

	ws.Route(ws.POST("/").
		To(h.RunJob).
		Doc("运行JobTask").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(pipeline.RunJobRequest{}).
		Writes(task.JobTask{}))

	ws.Route(ws.GET("/{id}").
		To(h.DescribeJobTask).
		Doc("查询JobTask运行任务详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(task.DescribeJobTaskRequest{}).
		Writes(task.JobTask{}))

	// 通过Job自身的Token进行认证
	ws.Route(ws.POST("/{id}/status").
		To(h.UpdateJobTaskStatus).
		Doc("更新任务状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Create.Value()).
		Reads(task.UpdateJobTaskStatusRequest{}).
		Writes(task.JobTask{}))

	// 通过Job自身的Token进行认证
	ws.Route(ws.POST("/{id}/output").
		To(h.UpdateJobTaskOutput).
		Doc("更新任务输出").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Create.Value()).
		Reads(task.UpdateJobTaskOutputRequest{}).
		Writes(task.JobTask{}))

	// Socket内鉴权
	ws.Route(ws.GET("/{id}/log").
		To(h.JobTaskLog).
		Doc("查询任务日志[WebSocket]").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Get.Value()).
		Reads(task.WatchJobTaskLogRequest{}).
		Writes(task.JobTaskStreamReponse{}))

	// Socket内鉴权
	ws.Route(ws.GET("/{id}/debug").
		To(h.JobTaskDebug).
		Doc("任务在线Debug[WebSocket]").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, JOB_TASK_RESOURCE_NAME).
		Metadata(label.Action, label.Get.Value()).
		Reads(task.JobTaskDebugRequest{}).
		Writes(task.JobTaskStreamReponse{}))
}

func (h *JobTaskHandler) QueryJobTask(r *restful.Request, w *restful.Response) {
	req := task.NewQueryTaskRequestFromHttp(r)
	set, err := h.service.QueryJobTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *JobTaskHandler) DescribeJobTask(r *restful.Request, w *restful.Response) {
	req := task.NewDescribeJobTaskRequest(r.PathParameter("id"))
	set, err := h.service.DescribeJobTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *JobTaskHandler) RunJob(r *restful.Request, w *restful.Response) {
	req := pipeline.NewRunJobRequest("")
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UpdateFromToken(token.GetTokenFromRequest(r))
	set, err := h.service.RunJob(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *JobTaskHandler) UpdateJobTaskOutput(r *restful.Request, w *restful.Response) {
	req := task.NewUpdateJobTaskOutputRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateJobTaskOutput(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *JobTaskHandler) UpdateJobTaskStatus(r *restful.Request, w *restful.Response) {
	req := task.NewUpdateJobTaskStatusRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateJobTaskStatus(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 60 * time.Second,
		ReadBufferSize:   8192,
		WriteBufferSize:  8192,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (h *JobTaskHandler) JobTaskLog(r *restful.Request, w *restful.Response) {
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}

	term := task.NewTaskLogWebsocketTerminal(ws)

	// 开启认证与鉴权
	entry := endpoint.NewEntryFromRestRequest(r)
	if entry != nil {
		entry.SetAuthEnable(true)
		entry.SetPermissionEnable(true)
		err = middleware.GetHttpAuther().PermissionCheck(r, w, entry)
		if err != nil {
			term.Failed(err)
			return
		}
	}

	// 读取请求
	in := task.NewWatchJobTaskLogRequest(r.PathParameter("id"))
	if err = term.ReadReq(in); err != nil {
		term.Failed(err)
		return
	}

	// 输出日志到Term中
	if err = h.service.WatchJobTaskLog(in, term); err != nil {
		term.Failed(err)
		return
	}

	term.Success("ok")
}

func (h *JobTaskHandler) JobTaskDebug(r *restful.Request, w *restful.Response) {
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}

	term := terminal.NewWebSocketTerminal(ws)

	// 开启认证与鉴权
	entry := endpoint.NewEntryFromRestRequest(r).
		SetAuthEnable(true).
		SetPermissionEnable(true)
	err = middleware.GetHttpAuther().PermissionCheck(r, w, entry)
	if err != nil {
		term.Failed(err)
		return
	}

	// 读取请求
	in := task.NewJobTaskDebugRequest(r.PathParameter("id"))
	if err = term.ReadReq(in); err != nil {
		term.Failed(err)
		return
	}

	// 进入容器
	in.SetWebTerminal(term)
	h.service.JobTaskDebug(r.Request.Context(), in)

	term.Success("ok")
}
