package callback

import "github.com/infraboard/mpaas/apps/approval"

func NewFeishuCardCallback() *FeishuCardCallback {
	return &FeishuCardCallback{
		Action: FeishuCardAction{},
	}
}

type FeishuCardAction struct {
	// 标签类型
	Tag string `json:"tag"`
	// 标签值
	Value map[string]string `json:"value"`
}

// 飞书卡片回调数据结构
type FeishuCardCallback struct {
	// 卡片操作用户的open_id
	OpenId string `json:"open_id"`
	// 卡片操作用户的user_id
	UserId string `json:"user_id"`
	// 卡片消息的唯一id
	OpenMessageId string `json:"open_message_id"`
	// 卡片消息归属的租户id
	TenantKey string `json:"tenant_key"`
	// 更新卡片用的token(凭证)
	Token string `json:"token"`
	// 卡片按钮携带的数据
	Action FeishuCardAction `json:"action"`
}

func (f *FeishuCardCallback) ApprovalId() string {
	return f.Action.Value["approval_id"]
}

func (f *FeishuCardCallback) Status() (approval.STAGE, error) {
	return approval.ParseSTAGEFromString(f.Action.Value["status"])
}

func (r *FeishuCardCallback) BuildUpdateApprovalStatusRequest() (*approval.UpdateApprovalStatusRequest, error) {
	req := approval.NewUpdateApprovalStatusRequest(r.ApprovalId())
	req.Status.AuditBy = ""

	stage, err := r.Status()
	if err != nil {
		return nil, err
	}
	req.Status.Stage = stage
	return req, nil
}
