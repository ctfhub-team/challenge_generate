package cmd

import (
	"cg/pkg/cmdutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ConfigCmd)
	ConfigCmd.AddCommand(ConfigSetCmd)
	ConfigCmd.AddCommand(ConfigGetCmd)
	ConfigCmd.AddCommand(ConfigCleanCmd)
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "设置 cg 的配置文件",
	Long:  "设置 cg 的配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var ConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: "进入配置设置向导",
	Long:  "进入配置设置向导",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigSet()
	},
}

var ConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取当前配置",
	Long:  "获取当前配置",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigGet()
	},
}

var ConfigCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清除当前配置",
	Long:  "清除当前配置",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigClean()
	},
}
