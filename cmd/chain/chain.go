package chain

import (
	"github.com/spf13/cobra"
)

func init() {
	CmdChain.AddCommand(cmdPs)
	CmdChain.AddCommand(cmdAdd)
	CmdChain.AddCommand(cmdRemove)
}

var CmdChain = &cobra.Command{
	Use:   "chain [COMMAND] [FLAGS]",
	Short: "隧道管理",
}
