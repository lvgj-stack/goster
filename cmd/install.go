package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdInstall = &cobra.Command{
	Use:   "install [URL]",
	Short: "下载并安装 gost",
	Long:  `下载并安装 gost，通过 url 可以指定下载的版本地址`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://github.com/go-gost/gost/releases/download/v3.0.0-rc10/gost_3.0.0-rc10_linux_amd64.tar.gz"
		if len(args) > 0 {
			url = args[0]
		}
		c := fmt.Sprintf("bash -c 'rm -rf goster_tmp "+
			"&& mkdir goster_tmp "+
			"&& cd goster_tmp "+
			"&& wget -O gost.tar.gz %s "+
			"&& tar -zxf ./gost.tar.gz"+
			"&& chmod +x ./gost "+
			"&& mv ./gost /usr/bin"+
			"&& cd .. "+
			"&& rm -rf ./goster_tmp'", url)
		err := common.Bash(c)
		if err != nil {
			common.Bash("rm -rf ./goster_tmp")
			return
		}

		err = common.Bash(`
mkdir -p /etc/gost && 
echo "{}" > /etc/gost/gost.json
`)
		if err != nil {
			return
		}

		err = common.Bash(`
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
ExecStart=/usr/bin/gost -C /etc/gost/gost.json

[Install]
WantedBy=multi-user.target
" > /lib/systemd/system/gost.service && systemctl enable gost.service && systemctl restart gost.service
`)
		if err != nil {
			return
		}
		fmt.Println("安装成功！")
	},
}
