package api

import (
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/trigger"
)

// 处理来自gitlab的事件
// Hook Header参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#delivery-headers
// 参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html
func (h *handler) HandleGitlabEvent(r *restful.Request, w *restful.Response) {
	event := trigger.NewGitlabWebHookEvent()
	event.ParseInfoFromHeader(r)
	err := r.ReadEntity(event)
	if err != nil {
		response.Failed(w, err)
	}

	req := trigger.NewGitlabEvent(event)
	req.Id = r.PathParameter(trigger.GITLAB_HEADER_EVENT_UUID)
	req.SkipRunPipeline, err = strconv.ParseBool(r.QueryParameter("skip_run_pipeline"))
	if err != nil {
		response.Failed(w, err)
	}

	h.log.Debugf("application %s accept event: %s", event)
	ins, err := h.svc.HandleEvent(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// 查询repo 的gitlab地址, 手动获取信息, 触发手动事件
func (h *handler) MannulGitlabEvent(r *restful.Request, w *restful.Response) {

}
