package api

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/trigger"
)

// 处理来自gitlab的事件
// Hook Header参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#delivery-headers
// 参考文档: https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html
func (h *Handler) HandleGitlabEvent(r *restful.Request, w *restful.Response) {
	event := trigger.NewGitlabWebHookEvent()
	event.ParseInfoFromHeader(r)

	// 读取body数据
	body, err := io.ReadAll(r.Request.Body)
	defer r.Request.Body.Close()
	if err != nil {
		response.Failed(w, err)
		return
	}
	event.EventRaw = string(body)

	// 反序列化
	err = json.Unmarshal(body, event)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := trigger.NewGitlabEvent(event)
	req.Id = r.PathParameter(trigger.GITLAB_HEADER_EVENT_UUID)
	req.SkipRunPipeline, err = strconv.ParseBool(r.QueryParameter("skip_run_pipeline"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	h.log.Debugf("accept event: %s", event)
	ins, err := h.svc.HandleEvent(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

// 查询repo 的gitlab地址, 手动获取信息, 触发手动事件
func (h *Handler) MannulGitlabEvent(r *restful.Request, w *restful.Response) {

}
