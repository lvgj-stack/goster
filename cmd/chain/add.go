package chain

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdAdd = &cobra.Command{
	Use:     "add [隧道名] [隧道协议] [隧道对端ip+端口] [可选：账号] [可选：密码]",
	Example: "goster chain add cmihk-01 relay+tls 193.55.55.2:40675 [username] [password]",
	Short:   "添加隧道",
	Args:    cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		var username, password string
		if len(args) >= 5 {
			username = args[3]
			password = args[4]
		}
		connector, dialer := args[1], args[1]
		protocolss := strings.Split(args[1], "+")
		if len(protocolss) == 2 {
			connector = protocolss[0]
			dialer = protocolss[1]
		}
		portSs := strings.Split(args[2], ":")
		port := portSs[len(portSs)-1]
		common.GlobalConfig.Chains = append(common.GlobalConfig.Chains, common.Chains{
			Name: args[0],
			Hops: []common.Hops{
				{
					Name: "hop-" + args[0],
					Nodes: []common.Nodes{
						{
							Name: "node-0",
							Addr: args[2],
							Connector: common.Connector{
								Type: connector,
								Auth: common.Auth{
									Username: username,
									Password: password,
								},
							},
							Dialer: common.Dialer{
								Type: dialer,
							},
						},
					},
				},
			},
		})
		changeBytes, err := json.Marshal(common.GlobalConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := os.WriteFile(common.ConfigPath, changeBytes, 0644); err != nil {
			fmt.Println(err.Error())
			return
		}
		agentConfig := fmt.Sprintf(`
{
  "services": [
    {
      "name": "service-0",
      "addr": ":%s",
      "handler": {
        "type": "%s",
		"auth": {
		  "username": "%s",
		  "password": "%s"
		},
        "metadata": {
          "bind": "true"
        }
      },
      "listener": {
        "type": "%s",
        "metadata": {
          "bind": "true"
        }
      },
	  "metadata": {
	    "bind": "true"
	  }
    }
  ]
}
`, port, connector, username, password, dialer)
		fmt.Println("-------------------------------------------")
		fmt.Println("如果对端没有安装Gost，请复制执行：")
		fmt.Printf(`
mkdir goster_tmp && 
cd goster_tmp &&
wget -O gost.tar.gz https://github.com/go-gost/gost/releases/download/v3.0.0-nightly.20240201/gost_3.0.0-nightly.20240201_linux_amd64.tar.gz &&
tar -zxf ./gost.tar.gz &&
chmod +x ./gost &&
mv ./gost /usr/bin &&
cd .. &&
rm -rf ./goster_tmp && 
mkdir -p /etc/gost && 
echo '%s' > %s &&
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
`, agentConfig, common.ConfigPath, common.ConfigPath)
		fmt.Println("-------------------------------------------")
		fmt.Println("如果对端已经安装Gost，请复制执行：")
		fmt.Printf(`
echo '%s' > %s && systemctl restart gost.service
`, agentConfig, common.ConfigPath)
		fmt.Println("-------------------------------------------")
		fmt.Println("|||  对端完成安装后，执行 goster restart  |||")
		fmt.Println("-------------------------------------------")
	},
}
