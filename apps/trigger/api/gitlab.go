package api

import (
	"context"
	"encoding/json"
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

// 查询repo 的gitlab地址, 手动获取信息, 触发手动事件
func (h *Handler) MannulGitlabEvent(r *restful.Request, w *restful.Response) {
	// 构造事件
	gevent := trigger.NewGitlabWebHookEvent()
	event := trigger.NewGitlabEvent(gevent)

	// 反序列化
	err := r.ReadEntity(event)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 事件关联信息填充
	err = h.BuildEvent(r.Request.Context(), event)
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

func (h *Handler) BuildEvent(ctx context.Context, in *trigger.Event) error {
	in.IsMannul = true

	// 查询服务仓库信息
	descReq := service.NewDescribeServiceRequest(in.GitlabEvent.EventToken)
	svc, err := h.mcenter.Service().DescribeService(ctx, descReq)
	if err != nil {
		return err
	}
	repo := svc.Spec.Repository

	// 补充Project相关信息
	p := in.GitlabEvent.Project
	p.Id = repo.ProjectIdToInt64()
	p.GitHttpUrl = repo.HttpUrl
	p.GitSshUrl = repo.SshUrl
	p.NamespacePath = repo.Namespace
	p.WebUrl = repo.WebUrl
	p.Name = svc.Spec.Name

	// 补充分支相关Commit信息
	gc, err := repo.MakeGitlabConfig()
	if err != nil {
		return err
	}
	v4 := gitlab.NewGitlabV4(gc)
	branchReq := gitlab.NewGetProjectBranchRequest()
	branchReq.ProjectId = repo.ProjectId
	branchReq.Branch = in.GitlabEvent.Ref
	b, err := v4.Project().GetProjectBranch(ctx, branchReq)
	if err != nil {
		return err
	}
	in.GitlabEvent.Commits = append(in.GitlabEvent.Commits, &trigger.Commit{
		Id:        b.Commit.Id,
		Message:   b.Commit.Message,
		Title:     b.Commit.Title,
		Timestamp: b.Commit.CommittedDate,
	})
	return nil
}
