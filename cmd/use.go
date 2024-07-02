package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdUse = &cobra.Command{
	Use:     "use [Gost配置路径]",
	Example: "goster use /etc/gost/config.json",
	Short:   "修改Gost默认使用的配置文件",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := &common.GosterConfig{
			ConfigName: args[0],
		}
		bs, err := json.Marshal(c)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := os.WriteFile(common.GosterConfigPath, bs, 0644); err != nil {
			fmt.Println(err.Error())
			return
		}
		newSystemdStr := fmt.Sprintf(`
echo "[Unit]
Description=gost
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
Type=simple
User=root
Restart=always
RestartSec=5
DynamicUser=true
ExecStart=/usr/bin/gost -C %s

[Install]
WantedBy=multi-user.target
" > /lib/systemd/system/gost.service && systemctl enable gost.service && systemctl restart gost.service
`, args[0])
		if err := common.Bash(newSystemdStr); err != nil {
			fmt.Errorf(err.Error())
			return
		}
	},
}
