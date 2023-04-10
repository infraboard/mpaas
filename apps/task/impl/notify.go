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
		return
	}

	status := task.NewCallbackStatus(mu.UserName)
	in.Status.AddNotifyStatus(status)

	// 调用mcenter api 通知用户
	for _, nt := range mu.NotifyTypes {
		switch nt {
		case notify.NOTIFY_TYPE_MAIL:
			req := notify.NewSendMailRequest(
				[]string{mu.UserName},
				in.ShowTitle(),
				in.HTMLContent(),
			)
			i.mcenter.Notify().SendMail(ctx, req)
		case notify.NOTIFY_TYPE_SMS:
			req := notify.NewSendSMSRequest()
			i.mcenter.Notify().SendSMS(ctx, req)
		case notify.NOTIFY_TYPE_VOICE:
			i.mcenter.Notify().SendVoice(ctx, nil)
		case notify.NOTIFY_TYPE_IM:
			i.mcenter.Notify().SendIM(ctx, nil)
		}
	}
}

// 调用mcenter api 通知用户Pipeline Task执行状态
func (i *impl) PipelineTaskMention(ctx context.Context, mu *pipeline.MentionUser, in *task.PipelineTask) {
	if !mu.IsMatch(in.Status.Stage.String()) {
		return
	}
}
