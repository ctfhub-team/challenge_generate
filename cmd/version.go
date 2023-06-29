package cmd

import (
	"cg/pkg/util"
	"fmt"

	"github.com/gookit/color"
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
		Cyan := color.FgCyan.Render
		fmt.Println("Version:   ", Cyan(util.Version))
		fmt.Println("BuildTime: ", Cyan(util.BuildTime))
		fmt.Println("GitCommit: ", Cyan(util.GitCommitId))
	},
}
