package impl

import (
	"context"

	"github.com/infraboard/mpaas/apps/approval"
)

// 创建发布申请
func (i *impl) CreateApproval(ctx context.Context, in *approval.CreateApprovalRequest) (
	*approval.Approval, error) {
	return nil, nil
}

// 查询发布申请列表
func (i *impl) QueryApproval(ctx context.Context, in *approval.QueryApprovalRequest) (
	*approval.ApprovalSet, error) {
	return nil, nil
}

// 查询发布申请详情
func (i *impl) DescribeApproval(ctx context.Context, in *approval.DescribeApprovalRequest) (
	*approval.Approval, error) {
	return nil, nil
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
