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
	Short: "进入设置向导",
	Long:  "进入设置向导",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigSet()
	},
}

var ConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "获取当前配置项",
	Long:  "获取当前配置项",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigGet()
	},
}
