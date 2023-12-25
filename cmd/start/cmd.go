package start

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/clients/rpc/hooks"

	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
	"github.com/infraboard/mcube/v2/ioc/server"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mpaas API服务",
	Long:  "mpaas API服务",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(server.SetUp(func() {
			hooks.NewMcenterAppHook().SetupAppHook()
		}).Run(context.Background()))
	},
}
