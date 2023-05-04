package notify

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/notify"
)

//go:embed templates/*
var templatesDir embed.FS

func NewFeishuAuditNotifyMessage() *FeishuAuditNotifyMessage {
	return &FeishuAuditNotifyMessage{
		ShowPassButton: true,
		ShowDenyButton: true,
	}
}

type FeishuAuditNotifyMessage struct {
	// 标题
	Title string
	// 申请人
	CreateBy string
	// 执行人
	Operator string
	// 其他审核人
	Auditor string
	// 流水线描述
	PipelineDesc string
	// 执行方式
	ExecType string
	// 执行时变量
	ExecVars string
	// 是否显示同意按钮
	ShowPassButton bool
	// 同意按钮的名称
	PassButton string
	// 是否显示拒绝按钮
	ShowDenyButton bool
	// 拒绝按钮的名称
	DenyButton string
	// 申请单Id, 点击按钮触发时的回调携带参数
	ApprovalId string
	// 备注
	Note string
}

func (t *FeishuAuditNotifyMessage) render() (string, error) {
	content, err := templatesDir.ReadFile("templates/feishu_card.tmpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("feishu_card").Parse(string(content))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, t)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (t *FeishuAuditNotifyMessage) BuildNotifyRequest(userIds ...string) (*notify.SendNotifyRequest, error) {
	if len(userIds) == 0 {
		return nil, fmt.Errorf("send feishu notify, but feishu user id is nil")
	}

	req := notify.NewSendNotifyRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.NotifyTye = notify.NOTIFY_TYPE_IM
	req.AddUser(userIds...)
	req.Title = t.Title
	req.ContentType = "interactive"

	content, err := t.render()
	if err != nil {
		return nil, err
	}
	req.Content = content
	return req, nil
}
