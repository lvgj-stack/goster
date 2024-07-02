package service

import (
	"github.com/spf13/cobra"
)

func init() {
	CmdService.AddCommand(cmdPs)
	CmdService.AddCommand(cmdAdd)
	CmdService.AddCommand(cmdRemove)
}

var CmdService = &cobra.Command{
	Use:   "service [COMMAND] [FLAGS]",
	Short: "转发管理",
}
