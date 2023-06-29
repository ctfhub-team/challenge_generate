package cmd

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cg",
	Short: "cg 是用于创建CTF题目环境模板的自助式生成工具",
	Long: `
CTF Generate (cg) 是用于创建CTF题目环境模板的自助式生成工具.

https://www.ctfhub.com/      Developed by: L1n3@CTFHub Team
`,
}

// func init() {
// 	RootCmd.CompletionOptions.DisableDefaultCmd = true
// }

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         RootCmd,
		Headings:        cc.HiGreen + cc.Underline,
		Commands:        cc.Cyan + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Cyan + cc.Bold,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
