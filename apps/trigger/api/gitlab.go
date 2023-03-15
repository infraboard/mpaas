package api

import (
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/trigger"
)

const (
	GitlabEventHeaderKey = "X-Gitlab-Event"
	GitlabEventTokenKey  = "X-Gitlab-Token"
)

// 处理来自gitlab的事件
// 参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html
func (h *handler) HandleGitlabEvent(r *restful.Request, w *restful.Response) {
	eventType := r.HeaderParameter(GitlabEventHeaderKey)
	serviceId := r.HeaderParameter(GitlabEventTokenKey)
	switch eventType {
	case "Push Hook":
		event := trigger.NewGitlabWebHookEvent()
		err := r.ReadEntity(event)
		if err != nil {
			response.Failed(w, err)
		}
		event.EventType = trigger.EVENT_TYPE_PUSH

		req := trigger.NewGitlabEvent(serviceId, event)
		req.SkipRunPipeline, err = strconv.ParseBool(r.QueryParameter("skip_run_pipeline"))
		if err != nil {
			response.Failed(w, err)
		}

		h.log.Debugf("application %s accept event: %s", serviceId, event)
		ins, err := h.svc.HandleEvent(r.Request.Context(), req)
		if err != nil {
			response.Failed(w, err)
			return
		}

		response.Success(w, ins)
	case "Tag Push Hook":
	case "Merge Request Hook":
	case "Note Hook":
	case "Issue Hook":
	}
}

// 查询repo 的gitlab地址, 手动获取信息, 触发手动事件
func (h *handler) MannulGitlabEvent(r *restful.Request, w *restful.Response) {

}
