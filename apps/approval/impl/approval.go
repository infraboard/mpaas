package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/approval"
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

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a approval document error, %s", err)
	}
	return nil, nil
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

	return ins, nil
}

// 编辑发布申请
func (i *impl) EditApproval(ctx context.Context, in *approval.EditApprovalRequest) (
	*approval.Approval, error) {
	return nil, nil
}

// 更新发布申请状态
func (i *impl) UpdateApprovalStatus(ctx context.Context, in *approval.UpdateApprovalStatusRequest) (
	*approval.Approval, error) {
	return nil, nil
}
