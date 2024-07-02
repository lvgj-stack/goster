package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdAdd = &cobra.Command{
	Use: "add [转发名] [协议(tcp,udp,tcp+udp,rtcp,rudp)] [本地端口] [目标IP+端口] [隧道名(可选)]",
	Example: `
新增转发时：需要填目标IP+端口
goster service add bestvm-hk tcp :40005 234.23.234.87:40098 chain-0
新增内网穿透时：只需要公网入口的端口，不需要ip；本地可以填ip
goster service add bestvm-hk tcp :40005 :40098 chain-0
`,
	Short: "添加转发规则",
	Args:  cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		var (
			chainName string
			handler   common.Handler
			listener  common.Listener
		)
		if len(args) == 5 {
			chainName = args[4]
		}
		if string(args[2][0]) != ":" && len(args[2]) <= 5 {
			args[2] = ":" + args[2]
		}
		addr, targetAddr := args[2], args[3]
		ss := strings.Split(args[1], "+")
		for _, protocol := range ss {
			handler.Type = protocol
			listener.Type = protocol
			if funk.ContainsString([]string{"rtcp", "rudp"}, protocol) {
				listener.Chain = chainName
				addr = args[3]
				targetAddr = args[2]
			} else {
				handler.Chain = chainName
			}
			common.GlobalConfig.Services = append(common.GlobalConfig.Services, common.Services{
				Name:     args[0] + "-" + protocol,
				Addr:     addr,
				Handler:  handler,
				Listener: listener,
				Forwarder: common.Forwarder{
					Nodes: []common.Nodes{
						{
							Name: "target-0",
							Addr: targetAddr,
						},
					},
				},
			})
		}
		changeBytes, err := json.Marshal(common.GlobalConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := os.WriteFile(common.ConfigPath, changeBytes, 0644); err != nil {
			fmt.Println(err.Error())
			return
		}
		common.Restart()
	},
}
