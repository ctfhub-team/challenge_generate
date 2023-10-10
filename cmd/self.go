package cmd

import (
	"cg/pkg/cmdutil"
	"os"

	"github.com/spf13/cobra"
)

var Proxy string

func init() {
	RootCmd.AddCommand(SelfCmd)
	SelfCmd.AddCommand(UpgradeCmd)
	SelfCmd.AddCommand(SelfCheckCmd)
	SelfCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "Socks5 Proxy example 1.2.3.4:5678")
}

var SelfCmd = &cobra.Command{
	Use:   "self",
	Short: "升级",
	Long:  "升级",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var UpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "自动升级",
	Long:  "自动升级",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Upgrade(Proxy)
	},
}

var SelfCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "检查新版本",
	Long:  "检查新版本",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.SelfCheck(Proxy)
	},
}
