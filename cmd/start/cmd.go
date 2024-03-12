package start

import (
	"github.com/spf13/cobra"

	"github.com/infraboard/mcube/v2/ioc/server"

	// 注册所有服务
	_ "github.com/infraboard/mpaas/apps"

	//
	_ "github.com/infraboard/mcenter/clients/rpc/middleware"
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mpaas API服务",
	Long:  "mpaas API服务",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(server.Run(cmd.Context()))
	},
}
