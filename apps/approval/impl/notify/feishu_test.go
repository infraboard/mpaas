package notify_test

import (
	"strings"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mpaas/apps/approval/impl/notify"
)

var feishuNotifyCard = &notify.FeishuAuditNotifyMessage{
	Domain:         domain.DEFAULT_DOMAIN,
	Namespace:      namespace.DEFAULT_NAMESPACE,
	Title:          "xxx提交的「系统帐号及权限申请」待你审批",
	CreateBy:       "[@王冰](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uUzN1YjL1cTN24SN3UjN?from=mcb)",
	Operator:       "xxxx",
	Auditor:        "xxx",
	PipelineDesc:   "xxx流水线",
	ExecType:       "手动执行",
	ExecVars:       "\nxx\nxxx\n",
	ShowPassButton: true,
	PassButtonName: "同意",
	ShowDenyButton: true,
	DenyButtonName: "拒绝",
	Note:           "备注信息",
	ApprovalId:     "xxx",
}

func TestFeishuCardAuditNotity(t *testing.T) {
	req, err := feishuNotifyCard.BuildNotifyRequest("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.Content)
}

func TestFeishuCardPassNotify(t *testing.T) {
	msg := feishuNotifyCard
	msg.ShowDenyButton = false
	msg.PassButtonName = "xxx已同意"
	req, err := feishuNotifyCard.BuildNotifyRequest("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.Content)
}

func TestFeishuCardDenyNotify(t *testing.T) {
	msg := feishuNotifyCard
	msg.ExecVars = strings.ReplaceAll(msg.ExecVars, "\n", "\\n")
	msg.ShowPassButton = false
	msg.DenyButtonName = "xxx已拒绝"
	req, err := msg.BuildNotifyRequest("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.Content)

	res, err := nrpc.SendNotify(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
