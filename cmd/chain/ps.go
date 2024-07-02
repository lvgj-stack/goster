package chain

import (
	"fmt"

	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdPs = &cobra.Command{
	Use:   "ps",
	Short: "查看所有隧道",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		if len(common.GlobalConfig.Chains) > 0 {
			fmt.Println("隧道:")
			chainTable, err := gotable.Create("备注/名称", "协议", "隧道地址")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, chain := range common.GlobalConfig.Chains {
				if len(chain.Hops) > 0 {
					chainTable.AddRow([]string{chain.Name, chain.Hops[0].Nodes[0].Connector.Type + "+" + chain.Hops[0].Nodes[0].Dialer.Type, chain.Hops[0].Nodes[0].Addr})
				}
			}
			fmt.Println(chainTable)
		}
	},
}
