package webhook

import (
	"context"
	"sync"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
)

func NewWebHook() *WebHook {
	return &WebHook{
		log: zap.L().Named("webhook"),
	}
}

type WebHook struct {
	log logger.Logger
}

func (h *WebHook) SendJobTaskStatus(ctx context.Context, hooks []*pipeline.WebHook, t *task.JobTask) {
	if t == nil {
		return
	}

	if len(hooks) == 0 {
		return
	}

	if len(hooks) > MAX_WEBHOOKS_PER_SEND {
		t.Status.AddErrorEvent("too many webhooks configs current: %d, max: %d", len(hooks), MAX_WEBHOOKS_PER_SEND)
		return
	}

	h.log.Debugf("start send task[%s] webhook, total %d", t.Spec.JobName, len(hooks))
	wg := &sync.WaitGroup{}
	for i := range hooks {
		req := newRequest(hooks[i], t, wg)
		go req.Push(ctx)
	}
	wg.Wait()
}

func (h *WebHook) SendPipelineTaskStatus(ctx context.Context, hooks []*pipeline.WebHook, t *task.PipelineTask) {
	if t == nil {
		return
	}

	if len(hooks) == 0 {
		return
	}

	// if len(hooks) > MAX_WEBHOOKS_PER_SEND {
	// 	t.Status.AddErrorEvent("too many webhooks configs current: %d, max: %d", len(hooks), MAX_WEBHOOKS_PER_SEND)
	// 	return
	// }

	// h.log.Debugf("start send task[%s] webhook, total %d", t.Spec.JobName, len(hooks))
	// wg := &sync.WaitGroup{}
	// for i := range hooks {
	// 	req := newRequest(hooks[i], t, wg)
	// 	go req.Push(ctx)
	// }
	// wg.Wait()
}
