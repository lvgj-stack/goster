package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/chain"
	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
	"code.byted.org/lvguojun.001/top-modify-tool/cmd/service"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&common.ConfigPath, "configPath", "c", "/etc/gost/gost.json", "配置文件路径")
	rootCmd.AddCommand(cmdInstall)
	rootCmd.AddCommand(cmdPs)
	rootCmd.AddCommand(chain.CmdChain)
	rootCmd.AddCommand(service.CmdService)
	rootCmd.AddCommand(cmdRestart)
	rootCmd.AddCommand(cmdStatus)
	rootCmd.AddCommand(cmdUse)
	rootCmd.AddCommand(cmdUninstall)
}

var rootCmd = &cobra.Command{
	Use:   "goster [COMMAND] [FLAG]",
	Short: "goster 是用于管理 Gost 转发的命令行工具，https://github.com/ginuerzh/gost",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
