package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the start command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "mpaas 服务初始化",
	Long:  "mpaas 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		return nil
	},
}

func init() {

	RootCmd.AddCommand(initCmd)
}
