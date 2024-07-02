package service

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
	Use:     "remove [转发名]",
	Example: "goster service remove service-0 service-1 service-2",
	Short:   "删除转发",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		common.InitConfig()
		newConfig := &common.GostConfig{}
		if err := copier.Copy(&newConfig.Chains, common.GlobalConfig.Chains); err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, service := range common.GlobalConfig.Services {
			if !funk.ContainsString(args, service.Name) {
				newConfig.Services = append(newConfig.Services, service)
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
		common.Restart()
	},
}
