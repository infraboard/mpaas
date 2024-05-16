package terminal_test

import (
	"testing"

	"github.com/infraboard/mpaas/common/terminal"
)

func TestGetCmdHandleFuncPing(t *testing.T) {
	ping := terminal.GetCmdHandleFunc("ping")
	resp := terminal.NewResponse()
	ping(terminal.NewRequest(), resp)
	t.Log(resp.ToJSON())
}
