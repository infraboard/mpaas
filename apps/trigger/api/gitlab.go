package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
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

func NewMannulGitlabEventRequest() *MannulGitlabEventRequest {
	return &MannulGitlabEventRequest{}
}

type MannulGitlabEventRequest struct {
	// 服务Id
	ServiceId string `json:"service_id"`
	// 分支
	Branch string `json:"branch"`
	// 是否触发Pipeline执行
	SkipRunPipeline bool `json:"skip_run_pipeline"`
}

// 查询repo 的gitlab地址, 手动获取信息, 触发手动事件
func (h *Handler) MannulGitlabEvent(r *restful.Request, w *restful.Response) {
	in := NewMannulGitlabEventRequest()

	// 反序列化
	err := r.ReadEntity(in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 构造事件
	event, err := h.MockEvent(r.Request.Context(), in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	h.log.Debugf("mannul event: %s", event)
	ins, err := h.svc.HandleEvent(r.Request.Context(), event)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *Handler) MockEvent(ctx context.Context, in *MannulGitlabEventRequest) (*trigger.Event, error) {
	// 查询服务仓库信息
	descReq := service.NewDescribeServiceRequest(in.ServiceId)
	svc, err := h.mcenter.Service().DescribeService(ctx, descReq)
	if err != nil {
		return nil, err
	}
	repo := svc.Spec.Repository

	// 查询分支信息
	gc, err := repo.MakeGitlabConfig()
	if err != nil {
		return nil, err
	}
	v4 := gitlab.NewGitlabV4(gc)
	branchReq := gitlab.NewGetProjectBranchRequest()
	branchReq.ProjectId = repo.ProjectId
	branchReq.Branch = in.Branch
	b, err := v4.Project().GetProjectBranch(ctx, branchReq)
	if err != nil {
		return nil, err
	}
	fmt.Println(b)

	// 构造事件
	gevent := trigger.NewGitlabWebHookEvent()
	event := trigger.NewGitlabEvent(gevent)
	event.IsMannul = true
	event.SkipRunPipeline = in.SkipRunPipeline
	event.Provider = trigger.EVENT_PROVIDER_GITLAB
	return event, nil
}
