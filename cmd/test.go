package cmd

import (
	"fmt"
	"os"

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
	Short:   "d 测试已完成的镜像",
	Long:    `d 测试已完成的镜像`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println()
		_, err := os.Stat("docker-compose.yml")
		if err == nil {
			fmt.Println("检测到当前目录存在 docker-compose.yml")
		} else if os.IsNotExist(err) {
			fmt.Println("当前目录不存在 docker-compose.yml，请跳转至题目的enviroment目录后再进行测试")
		}
		os.Exit(0)
	},
}

var AutoCmd = &cobra.Command{
	Use:     "auto",
	Short:   "a 依次执行 Stop -> Build -> Run -> Bash",
	Long:    `a 构建镜像`,
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "b 构建镜像",
	Long:    `b 构建镜像`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Printf("构建镜像")
		os.Exit(0)
	},
}

var RunCmd = &cobra.Command{
	Use:     "run",
	Short:   "r 启动镜像",
	Long:    `r 启动镜像`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var StopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "s 停止镜像",
	Long:    `s 停止镜像`,
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
