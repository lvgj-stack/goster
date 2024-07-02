package cmd

import (
	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载gost、goster相关文件",
	Run: func(cmd *cobra.Command, args []string) {
		common.Bash(`
systemctl disable gost.service &&
rm -rf /etc/gost &&
rm -f /usr/bin/gost &&
rm -f /lib/systemd/system/gost.service &&
rm -f /usr/bin/goster
`)
	},
}
