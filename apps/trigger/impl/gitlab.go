package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/trigger"
)

// 应用事件处理
func (s *impl) HandleGitlabEvent(ctx context.Context, in *trigger.GitlabWebHookEvent) (
	*trigger.GitlabWebHookEvent, error) {
	return nil, nil
}
