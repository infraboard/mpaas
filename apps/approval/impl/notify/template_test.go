package notify_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/approval/impl/notify"
)

func TestFeishuAuditNotifyTemplate(t *testing.T) {
	msg := notify.NewFeishuAuditNotifyMessage()
	content, err := msg.Render()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}
