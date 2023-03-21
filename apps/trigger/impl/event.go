package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/trigger"
)

// 应用事件处理
func (i *impl) HandleEvent(ctx context.Context, in *trigger.Event) (
	*trigger.Record, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := trigger.NewRecord(in)

	switch in.Provider {
	case trigger.EVENT_PROVIDER_GITLAB:
		// 校验请求
		if err := in.GitlabEvent.Validate(); err != nil {
			return nil, exception.NewBadRequest(err.Error())
		}

		// 获取该服务对应事件的构建配置
		req := build.NewQueryBuildConfigRequest()
		req.AddService(in.GitlabEvent.ServiceId)
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

			bs := trigger.NewBuildStatus(buildConf)
			runReq := pipeline.NewRunPipelineRequest(pipelineId)
			runReq.RunBy = "gitlab_trigger"
			runReq.DryRun = in.SkipRunPipeline
			// 补充Git信息
			runReq.AddRunParam(in.GitlabEvent.GitRunParams()...)
			// 补充版本信息
			switch buildConf.Spec.VersionNamedRule {
			case build.VERSION_NAMED_RULE_DATE_BRANCH_COMMIT:
				runReq.AddRunParam(in.GitlabEvent.VersionRunParam(buildConf.Spec.VersionPrefix))
			}

			pt, err := i.task.RunPipeline(ctx, runReq)
			if err != nil {
				bs.ErrorMessage = err.Error()
			} else {
				bs.PiplineTaskId = pt.Meta.Id
				bs.PiplineTask = pt
			}

			ins.AddBuildStatus(bs)
		}
	}

	// 保存
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a deploy document error, %s", err)
	}
	return ins, nil
}

// 查询事件
func (i *impl) QueryRecord(ctx context.Context, in *trigger.QueryRecordRequest) (
	*trigger.RecordSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find event record error, error is %s", err)
	}

	set := trigger.NewRecordSet()
	// 循环
	for resp.Next(ctx) {
		ins := trigger.NewDefaultRecord()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode event record error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get event record count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
