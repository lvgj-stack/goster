package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	strs []string
	it   int
	str  string
)

func init() {
	modifyUsersQps.Flags().StringVarP(&str, "str", "o", "./ecs_new_api.json", "示例str")
	modifyUsersQps.Flags().StringSliceVarP(&strs, "strs", "a", []string{}, "示例strs")
	modifyUsersQps.Flags().IntVarP(&it, "it", "q", 0, "示例int")
}

var modifyUsersQps = &cobra.Command{
	Use:   "example",
	Short: "示例",
	Long:  `示例`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args, strs, str, it)
	},
}
