package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/trigger"
)

// 应用事件处理
func (i *impl) HandleServiceEvent(ctx context.Context, in *trigger.ServiceEvent) (
	*trigger.ServiceEvent, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	switch in.Provider {
	case trigger.EVENT_PROVIDER_GITLAB:
		// 校验请求
		if err := in.GitlabEvent.Validate(); err != nil {
			return nil, exception.NewBadRequest(err.Error())
		}

		// 获取该服务对应事件的构建配置
		req := build.NewQueryBuildConfigRequest()
		req.AddService(in.ServiceId)
		req.Event = in.GitlabEvent.EventName
		set, err := i.build.QueryBuildConfig(ctx, req)
		if err != nil {
			return nil, err
		}

		matched := set.MatchBranch(in.GitlabEvent.GetBranche())
		for index := range matched.Items {
			// 执行构建配置匹配的流水线
			buildConf := matched.Items[index]
			pipelineId := buildConf.Spec.PipielineId()
			if pipelineId == "" {
				i.log.Debugf("构建配置: %s, 未配置流水线", buildConf.Spec.Name)
				continue
			}

			i.task.RunPipeline(ctx, task.NewRunPipelineRequest(pipelineId))
		}
	}

	return in, nil
}
