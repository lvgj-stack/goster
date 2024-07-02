package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdRestart = &cobra.Command{
	Use:   "restart",
	Short: "重启Gost",
	Run: func(cmd *cobra.Command, args []string) {
		if err := common.Bash("systemctl restart gost.service"); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
