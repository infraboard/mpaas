package notify_test

import (
	"testing"

	"github.com/infraboard/mpaas/apps/approval/impl/notify"
)

func TestFeishuAuditNotifyTemplate(t *testing.T) {
	temp, err := notify.FeishuAuditNotifyTemplate()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}
