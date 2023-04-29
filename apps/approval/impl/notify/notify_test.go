package notify_test

import (
	"context"

	"github.com/infraboard/mpaas/test/tools"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/client/rpc"
)

var (
	nrpc notify.RPCClient
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	nrpc = rpc.C().Notify()
}
