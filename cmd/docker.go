package cmd

import (
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
}

var DockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "测试已完成的环境",
	Long:    `测试已完成的环境`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println()
		os.Chdir("enviroment")
		Cyan := color.FgCyan.Render
		Red := color.FgRed.Render
		_, err := os.Stat("docker-compose.yml")
		if err == nil {
			fmt.Println(Cyan("检测到环境目录存在 docker-compose.yml"))
		} else if os.IsNotExist(err) {
			fmt.Println(Red("环境目录不存在 docker-compose.yml"))
		}
		os.Exit(0)
	},
}

var AutoCmd = &cobra.Command{
	Use:     "auto",
	Short:   "依次执行 Stop -> Build -> Run -> Bash",
	Long:    `构建镜像`,
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "构建镜像",
	Long:    `构建镜像`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Printf("开始构建镜像")
		os.Exit(0)
	},
}

var RunCmd = &cobra.Command{
	Use:     "run",
	Short:   "启动镜像",
	Long:    `启动镜像`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var StopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "停止镜像",
	Long:    `停止镜像`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var BashCmd = &cobra.Command{
	Use:   "bash",
	Short: "执行bash进入容器",
	Long:  `执行bash进入容器`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}
