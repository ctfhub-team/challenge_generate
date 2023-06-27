package cmd

import (
	"cg/pkg/cmdutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.AddCommand(WizardCmd)
	newCmd.AddCommand(FileCmd)
}

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "n 创建新的题目模板",
	Long:    `n 创建新的题目模板`,
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var WizardCmd = &cobra.Command{
	Use:     "wizard",
	Short:   "w 使用向导创建",
	Long:    `使用向导创建`,
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Wizard()
	},
}

var FileCmd = &cobra.Command{
	Use:     "file",
	Short:   "f 从预定义文件创建",
	Long:    `从预定义文件创建`,
	Aliases: []string{"f"},
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.File()
	},
}
