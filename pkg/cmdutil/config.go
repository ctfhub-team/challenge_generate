package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/tpl"
	"cg/pkg/util"
	"fmt"
	"os"

	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

func ConfigSet() {
	config := global.Config{}
	_ = yaml.Unmarshal(tpl.Config, &config)
	// 修改内容
	config.Author = util.InputLine("请输入你的ID: ")
	config.Contact = util.InputLine("请输入你的邮箱: ")
	config.RegistryUrl = util.SelectOne("请选择您要使用的默认镜像源", global.Registry)
	// 写入文件
	writeData, _ := yaml.Marshal(&config)
	UserHomeDir, _ := os.UserHomeDir()
	os.MkdirAll(UserHomeDir+"/.config/cg/", os.ModePerm)
	util.WriteFile(UserHomeDir+"/.config/cg/config.yaml", string(writeData), 0644)
}

func ConfigGet() {
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, err := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		fmt.Println("请确保已经设置配置")
		return
	}
	_ = yaml.Unmarshal(data, &config)
	Cyan := color.FgCyan.Render
	fmt.Println("制作者 ID: " + Cyan(config.Author))
	fmt.Println("制作者邮箱: ", Cyan(config.Contact))
	fmt.Println("默认镜像源地址: ", Cyan(config.RegistryUrl))
}

func ConfigClean() {
	UserHomeDir, _ := os.UserHomeDir()
	err := os.Remove(UserHomeDir + "/.config/cg/config.yaml")
	if err != nil {
		fmt.Println("清除错误: ", err)
		return
	}
	fmt.Println("配置已清除")
}
