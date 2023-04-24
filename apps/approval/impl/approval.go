package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/apps/pipeline"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 创建发布申请
func (i *impl) CreateApproval(ctx context.Context, in *approval.CreateApprovalRequest) (
	*approval.Approval, error) {
	ins, err := approval.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 检查user的有效性
	if err := i.CheckApprovalUser(ctx, ins); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 补充Pipeline创建
	if !in.IsTemplate && ins.Spec.PipelineId == "" {
		p, err := i.pipeline.CreatePipeline(ctx, in.PipelineSpec)
		if err != nil {
			return nil, err
		}

		ins.Spec.PipelineId = p.Meta.Id
	}

	// 保存申请单
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a approval document error, %s", err)
	}
	return ins, nil
}

func (i *impl) CheckApprovalUser(ctx context.Context, ins *approval.Approval) error {
	req := user.NewQueryUserRequest()
	req.UserIds = ins.Spec.UserIds()
	set, err := i.mcenter.User().QueryUser(ctx, req)
	if err != nil {
		return err
	}
	for _, uid := range req.UserIds {
		if !set.HasUser(uid) {
			return fmt.Errorf("uid %s not found", uid)
		}
	}
	return nil
}

// 查询发布申请列表
func (i *impl) QueryApproval(ctx context.Context, in *approval.QueryApprovalRequest) (
	*approval.ApprovalSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := approval.NewApprovalSet()
	// 循环
	for resp.Next(ctx) {
		ins := approval.NewDefaultApproval()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode deploy error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

// 查询发布申请详情
func (i *impl) DescribeApproval(ctx context.Context, in *approval.DescribeApprovalRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := approval.NewDefaultApproval()
	if err := i.col.FindOne(ctx, bson.M{"_id": in.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("approval %s not found", in)
		}

		return nil, exception.NewInternalServerError("find approval %s error, %s", in.Id, err)
	}

	p, err := i.pipeline.DescribePipeline(ctx, pipeline.NewDescribePipelineRequest(ins.Spec.PipelineId))
	if err != nil {
		return nil, err
	}
	ins.Pipeline = p

	return ins, nil
}

// 编辑发布申请
func (i *impl) EditApproval(ctx context.Context, in *approval.EditApprovalRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := approval.NewDescribeApprovalRequest(in.Id)
	ins, err := i.DescribeApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	// 1. 只有处于草稿状态的申请才允许编辑
	if !ins.Status.Stage.Equal(approval.STAGE_DRAFT) {
		if err != nil {
			return nil, exception.NewBadRequest("只有处于草稿状态的发布申请才能编辑")
		}
	}

	switch in.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Spec = in.Spec
	case request.UpdateMode_PATCH:
		if err := mergo.MergeWithOverwrite(ins.Spec, in.Spec); err != nil {
			return nil, err
		}
		if err := ins.Spec.Validate(); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", in.UpdateMode)
	}

	// 校验更新后请求合法性
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	ins.Meta.UpdateAt = time.Now().Unix()
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": ins})
	if err != nil {
		return nil, exception.NewInternalServerError("update approval(%s) error, %s", ins.Meta.Id, err)
	}

	return ins, nil
}

// 更新发布申请状态
func (i *impl) UpdateApprovalStatus(ctx context.Context, in *approval.UpdateApprovalStatusRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := approval.NewDescribeApprovalRequest(in.Id)
	ins, err := i.DescribeApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	// 1. 关闭后的发布申请 不能修改状态
	if !ins.Status.Stage.Equal(approval.STAGE_CLOSED) {
		if err != nil {
			return nil, exception.NewBadRequest("发布申请已关闭, 禁止更新状态")
		}
	}

	// 2. 修改的状态不能回退, 比如你不能把发布中的状态 修改为审核中
	if in.Status.Stage < ins.Status.Stage {
		return nil, exception.NewBadRequest("不能回退状态, 当前状态: %s", ins.Status.Stage)
	}

	// 3. 只有审核人能修改审核状态
	if in.Status.Stage.Equal(approval.STAGE_PENDDING) && !ins.Spec.IsAuditor(in.UpdateBy) {
		return nil, exception.NewBadRequest("只有审核人员: %s 能审核", ins.Spec.Auditors)
	}

	// 4. 保存更新
	ins.Status.Update(in.Status.Stage)
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": bson.M{"status": in.Status}})
	if err != nil {
		return nil, exception.NewInternalServerError("update approval(%s) error, %s", ins.Meta.Id, err)
	}

	// 5. 如果允许自动执行, 则审核通过后自动执行
	if ins.Spec.AutoRun && ins.Status.Stage.Equal(approval.STAGE_PASSED) {
		runReq := pipeline.NewRunPipelineRequest(ins.Spec.PipelineId)
		runReq.RunBy = "@" + ins.UUID()
		runReq.TriggerMode = pipeline.TRIGGER_MODE_APPROVAL
		runReq.AddRunParam(ins.Spec.RunParams...)
		runReq.ApprovalId = ins.Meta.Id
		pt, err := i.task.RunPipeline(ctx, runReq)
		if err != nil {
			return nil, err
		}
		i.log.Debugf("auto publish pipeline task: %s", pt.Meta.Id)
	}
	return ins, nil
}

// 删除发布申请
func (i *impl) DeleteApproval(ctx context.Context, in *approval.DeleteApprovalRequest) (
	*approval.Approval, error) {
	ins, err := i.DescribeApproval(ctx, approval.NewDescribeApprovalRequest(in.Id))
	if err != nil {
		return nil, err
	}

	// 未关闭的申请不允许删除
	if !ins.Status.Stage.Equal(approval.STAGE_CLOSED) {
		return nil, exception.NewBadRequest("申请单未关闭")
	}

	// 删除Pipeline
	_, err = i.pipeline.DeletePipeline(ctx, pipeline.NewDeletePipelineRequest(ins.Spec.PipelineId))
	if err != nil {
		return nil, err
	}

	// 删除Pipeline
	_, err = i.col.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete approval(%s) error, %s", in.Id, err)
	}

	return ins, nil
}
