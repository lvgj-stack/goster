package chain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/copier"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	"code.byted.org/lvguojun.001/top-modify-tool/cmd/common"
)

var cmdRemove = &cobra.Command{
	Use:     "remove [隧道名]",
	Example: "goster chain remove cmihk-01 cmihk-02 cmihk-03",
	Short:   "添加隧道",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		newConfig := &common.GostConfig{}
		if err := copier.Copy(&newConfig.Services, common.GlobalConfig.Services); err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, chain := range common.GlobalConfig.Chains {
			if !funk.ContainsString(args, chain.Name) {
				newConfig.Chains = append(newConfig.Chains, chain)
			}
		}
		changeBytes, err := json.Marshal(newConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := os.WriteFile(common.ConfigPath, changeBytes, 0644); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
