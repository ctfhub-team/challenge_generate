package cmd

import (
	"cg/pkg/cmdutil"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DockerCmd)
	DockerCmd.AddCommand(AutoCmd)
	DockerCmd.AddCommand(BuildCmd)
	DockerCmd.AddCommand(RunCmd)
	DockerCmd.AddCommand(StopCmd)
	DockerCmd.AddCommand(BashCmd)
	DockerCmd.AddCommand(LogCmd)
	DockerCmd.AddCommand(SaveCmd)
}

var Cyan = color.FgCyan.Render
var Red = color.FgRed.Render

func CheckDockerCompose() bool {
	_, err := os.Stat("docker-compose.yml")
	if err == nil {
		fmt.Println(Cyan("检测到当前目录存在 docker-compose.yml"))
	} else if os.IsNotExist(err) {
		fmt.Println(Red("当前目录不存在 docker-compose.yml"))
	}
	return err == nil
}

var DockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "docker相关操作",
	Long:    `docker相关操作`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println()
		CheckDockerCompose()
		os.Exit(0)
	},
}

var AutoCmd = &cobra.Command{
	Use:     "auto",
	Short:   "依次执行 Stop -> Build -> Run -> Bash",
	Long:    `构建镜像`,
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始测试")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Auto()
	},
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "构建镜像",
	Long:    `构建镜像`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始构建镜像")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Build()
	},
}

var RunCmd = &cobra.Command{
	Use:     "run",
	Short:   "运行镜像",
	Long:    `运行镜像`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始运行镜像")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Run()
	},
}

var StopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "停止镜像",
	Long:    `停止镜像`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始停止镜像")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Stop()
	},
}

var BashCmd = &cobra.Command{
	Use:   "bash",
	Short: "执行bash进入容器",
	Long:  `执行bash进入容器`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始进入容器")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Bash()
	},
}

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "查看容器日志",
	Long:  `查看容器日志`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("查看容器日志")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Log()
	},
}

var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "导出容器tar包",
	Long:  `导出容器tar包`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("导出容器tar包")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Save()
	},
}
