package notify_test

import (
	"strings"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	mcenter_notify "github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mpaas/apps/approval/impl/notify"
)

var feishuNotifyCard = &notify.FeishuAuditNotifyMessage{
	Title:          "xxx提交的「系统帐号及权限申请」待你审批",
	CreateBy:       "[@王冰](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uUzN1YjL1cTN24SN3UjN?from=mcb)",
	Operator:       "xxxx",
	Auditor:        "xxx",
	PipelineDesc:   "xxx流水线",
	ExecType:       "手动执行",
	ExecVars:       "\nxx\nxxx\n",
	ShowPassButton: true,
	PassButton:     "同意",
	ShowDenyButton: true,
	DenyButton:     "拒绝",
	Note:           "备注信息",
	ApprovalId:     "xxx",
}

func TestFeishuCardAuditNotity(t *testing.T) {
	msg := feishuNotifyCard
	content, err := msg.Render()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}

func TestFeishuCardPassNotify(t *testing.T) {
	msg := feishuNotifyCard
	msg.ShowDenyButton = false
	msg.PassButton = "xxx已同意"
	content, err := msg.Render()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}

func TestFeishuCardDenyNotify(t *testing.T) {
	msg := feishuNotifyCard
	msg.ExecVars = strings.ReplaceAll(msg.ExecVars, "\n", "\\n")
	msg.ShowPassButton = false
	msg.DenyButton = "xxx已拒绝"
	content, err := msg.Render()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)

	req := mcenter_notify.NewSendNotifyRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.NotifyTye = mcenter_notify.NOTIFY_TYPE_IM
	req.AddUser("admin")
	req.Title = msg.Title
	req.ContentType = "interactive"
	req.Content = content
	res, err := nrpc.SendNotify(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
