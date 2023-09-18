package cmd

import (
	"cg/pkg/cmdutil"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CheckCmd)
}

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "检查当前题目目录及内容是否符合规范",
	Long:  "检查当前题目目录及内容是否符合规范",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Check()
	},
}
