package service

import (
	"fmt"

	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdPs = &cobra.Command{
	Use:   "ps",
	Short: "查看所有转发规则",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		if len(common.GlobalConfig.Services) > 0 {
			forwardTable, err := gotable.Create("备注/名称", "协议", "本地端口", "目的端口", "隧道")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			penetrateTable, err := gotable.Create("备注/名称", "协议", "内网穿透的本地端口", "对端入口端口", "隧道")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, service := range common.GlobalConfig.Services {
				if !funk.ContainsString([]string{"rtcp", "rudp"}, service.Listener.Type) {
					if len(service.Forwarder.Nodes) > 0 {
						forwardTable.AddRow([]string{service.Name, service.Listener.Type, service.Addr, service.Forwarder.Nodes[0].Addr, service.Handler.Chain})
					} else {
						forwardTable.AddRow([]string{service.Name, service.Listener.Type, service.Addr, "", service.Handler.Chain})
					}
				} else {
					if len(service.Forwarder.Nodes) > 0 {
						penetrateTable.AddRow([]string{service.Name, service.Listener.Type, service.Forwarder.Nodes[0].Addr, service.Addr, service.Listener.Chain})
					} else {
						penetrateTable.AddRow([]string{service.Name, service.Listener.Type, "", service.Addr, service.Listener.Chain})
					}
				}
			}
			fmt.Println("转发:")
			fmt.Println(forwardTable)
			fmt.Println("内网穿透:")
			fmt.Println(penetrateTable)
		}
	},
}
