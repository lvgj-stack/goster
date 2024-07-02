package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "查看Gost状态",
	Run: func(cmd *cobra.Command, args []string) {
		if err := common.Bash("systemctl status gost.service"); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
