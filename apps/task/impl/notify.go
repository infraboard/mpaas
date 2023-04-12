package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
)

// 调用mcenter api 通知用户Job Task执行状态
func (i *impl) JotTaskMention(ctx context.Context, mu *pipeline.MentionUser, in *task.JobTask) {
	if !mu.IsMatch(in.Status.Stage.String()) {
		i.log.Debugf("stage: %s not matched target: %s", mu.Events, in.Status.Stage)
		return
	}

	status := task.NewCallbackStatus(mu.UserName)
	// 调用mcenter api 通知用户
	for _, nt := range mu.NotifyTypes {
		switch nt {
		case notify.NOTIFY_TYPE_MAIL:
			req := notify.NewSendMailRequest(
				in.ShowTitle(),
				in.HTMLContent(),
				mu.UserName,
			)
			resp, err := i.mcenter.Notify().SendNotify(ctx, req)
			if err != nil {
				status.AddEvent(task.NewErrorEvent(err.Error()))
			} else {
				status.AddEvent(task.NewDebugEvent(resp.ToJson()))
				message := resp.FailedResponseToMessage()
				if message != "" {
					status.AddEvent(task.NewErrorEvent(message))
				}
			}
		case notify.NOTIFY_TYPE_SMS:
			status.AddEvent(task.NewErrorEvent("sms not impl"))
		case notify.NOTIFY_TYPE_VOICE:
			status.AddEvent(task.NewErrorEvent("voice not impl"))
		case notify.NOTIFY_TYPE_IM:
			req := notify.NewSendNotifyRequest()
			req.Domain = in.Spec.Domain
			req.Namespace = in.Spec.Namespace
			req.NotifyTye = notify.NOTIFY_TYPE_IM
			req.AddUser(mu.UserName)
			req.Title = in.ShowTitle()
			req.Content = in.MarkdownContent()
			resp, err := i.mcenter.Notify().SendNotify(ctx, req)
			if err != nil {
				status.AddEvent(task.NewErrorEvent(err.Error()))
			} else {
				status.AddEvent(task.NewDebugEvent(resp.ToJson()))
				message := resp.FailedResponseToMessage()
				if message != "" {
					status.AddEvent(task.NewErrorEvent(message))
				}
			}
		}
	}
	status.MakeStatusUseEvent()
	in.Status.AddNotifyStatus(status)
}

// 调用mcenter api 通知用户Pipeline Task执行状态
func (i *impl) PipelineTaskMention(ctx context.Context, mu *pipeline.MentionUser, in *task.PipelineTask) {
	if !mu.IsMatch(in.Status.Stage.String()) {
		return
	}
}
