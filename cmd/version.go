package cmd

import (
	"cg/pkg/util"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出 cg 的版本和更新时间",
	Long:  "输出 cg 的版本和更新时间",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:    ", util.Version)
		fmt.Println("UpdateTime: ", util.UpdateTime)
	},
}
