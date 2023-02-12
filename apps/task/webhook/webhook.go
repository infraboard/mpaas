package webhook

import (
	"context"
	"fmt"

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

func (h *WebHook) Send(ctx context.Context, hooks []*pipeline.WebHook, t *task.JobTask) error {
	if t == nil {
		return fmt.Errorf("task is nil")
	}

	if err := h.validate(hooks); err != nil {
		return err
	}

	h.log.Debugf("start send task[%s] webhook, total %d", t.Spec.JobName, len(hooks))
	for i := range hooks {
		req := newRequest(hooks[i], t)
		req.Push()
	}

	return nil
}

func (h *WebHook) validate(hooks []*pipeline.WebHook) error {
	if len(hooks) == 0 {
		return nil
	}

	if len(hooks) > MAX_WEBHOOKS_PER_SEND {
		return fmt.Errorf("too many webhooks configs current: %d, max: %d", len(hooks), MAX_WEBHOOKS_PER_SEND)
	}

	return nil
}
