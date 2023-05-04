package notify

import (
	"bytes"
	"embed"
	"text/template"

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

// 关于飞书卡片数据结构描述: https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN
// 关于飞书卡片搭建工具: https://open.feishu.cn/document/ukTMukTMukTM/uYzM3QjL2MzN04iNzcDN/message-card-builder
type FeishuAuditNotifyMessage struct {
	// 标题
	Title string
	// 申请人
	CreateBy string
	// 执行人列表
	Operator string
	// 审核人列表
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
	PassButtonName string
	// 是否显示拒绝按钮
	ShowDenyButton bool
	// 拒绝按钮的名称
	DenyButtonName string
	// 申请单Id, 点击按钮触发时的回调携带参数
	ApprovalId string
	// 备注
	Note string

	// 需要关联传递的参数
	// 消息来源域
	Domain string
	// 消息来源空间
	Namespace string
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

func (t *FeishuAuditNotifyMessage) BuildNotifyRequest() (*notify.SendNotifyRequest, error) {
	req := notify.NewSendNotifyRequest()
	req.Domain = t.Domain
	req.Namespace = t.Namespace
	req.NotifyTye = notify.NOTIFY_TYPE_IM
	req.Title = t.Title
	req.ContentType = "interactive"

	content, err := t.render()
	if err != nil {
		return nil, err
	}
	req.Content = content
	return req, nil
}
