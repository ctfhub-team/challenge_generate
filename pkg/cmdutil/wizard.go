package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/util"
	"fmt"
	"os"

	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

/**
 * @Description: 生成题目模板
 * @param baseImageName 题目基础镜像名称
 * @param challengeType 题目类型
 * @param challengeName 题目名称
 * @param language 题目语言
 * @param hasDB 是否需要数据库
 */
func Generate(challengeInfo map[string]string) {
	// flag位置确认
	challengeInfo["need_flag"] = util.SelectArray("是否需要单独处理flag位置(flag.sh)", []string{"否", "是"})
	challengeInfo["need_start"] = util.SelectArray("是否需要单独处理部分服务启动(start.sh)", []string{"否", "是"})
	// 题目等级选择
	challengeInfo["level"] = util.SelectArray("此题目难度为", global.Level)

	// // DEBUG 测试输出
	// s, _ := json.MarshalIndent(challengeInfo, "", "    ")
	// fmt.Printf(string(s))

	// 二次确认
	confirm := util.SelectArray("确认创建题目 "+challengeInfo["challenge_name"]+" ?", []string{"确认", "取消"})
	if confirm == "取消" {
		os.Exit(0)
	}

	// 创建题目模板目录
	os.Mkdir(challengeInfo["challenge_name"], 0755)
	os.Chdir(challengeInfo["challenge_name"])

	// 创建文件夹
	dirTree := []string{
		"environment/src/",
		"environment/files/",
		"writeup/",
	}
	for _, path := range dirTree {
		os.MkdirAll(path, os.ModePerm)
	}
	// 写入dockerfile
	GenerateDockerFile(challengeInfo)

	// 写入docker-compose.yml
	GenerateDockerCompose(challengeInfo)

	// 写入meta.yml
	GenerateMeta(challengeInfo)

	// 写入数据库
	GenerateDB(challengeInfo)

	// 写入flag.sh处理
	GenerateFlag(challengeInfo)

	// 写入start.sh处理
	GenerateStart(challengeInfo)

	// 写入README.md
	GenerateReadme(challengeInfo)

	// 输出成功提示
	fmt.Println("")
	fmt.Println(challengeInfo["challenge_name"] + " 创建成功，请按如下步骤依次操作：")
	fmt.Println("1. 初始化Git仓库")
	fmt.Println("2. 编辑 " + challengeInfo["challenge_name"] + "/meta.yml 文件修改题目信息")
	fmt.Println("3. 进入 " + challengeInfo["challenge_name"] + "/environment 目录进行测试")
	fmt.Println("4. 进入 " + challengeInfo["challenge_name"] + "/writeup 目录补全WP及Exp")
	fmt.Println("5. 一切完成后推送至Git仓库")
}

func Wizard() {
	challengeInfo := map[string]string{
		"type":             "", // 题目类型
		"language":         "", // 使用语言
		"language_version": "", // 语言版本，HTML题目留空
		"webserver":        "", // Web服务器，非Web题目留空
		"db":               "", // 数据库，无数据库留空
		"pwn_arch":         "", // Pwn题目架构，非Pwn题目留空
		"pwn_server":       "", // Pwn题目服务器，非Pwn题目留空
		"need_flag":        "", // 是否需要flag.sh
		"need_start":       "", // 是否需要start.sh
		"level":            "", // 题目等级
		"base_image_name":  "", // 基础镜像名称
		"base_registry":    "", // 基础镜像源地址
		"challenge_name":   "", // 题目镜像名称
	}
	// 判断配置中是否已设置默认镜像源
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, err := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
	if err != nil {
		fmt.Println("未检测到配置文件，建议先设置默认镜像源")
		registry := util.SelectOne("请选择您要使用的镜像源", global.Registry)
		challengeInfo["base_registry"] = registry + "/"
	} else {
		_ = yaml.Unmarshal(data, &config)
		fmt.Println("检测到配置文件，将使用配置文件中的镜像源")
		fmt.Println("镜像源地址：" + color.FgCyan.Render(config.RegistryUrl))
		fmt.Println()
		challengeInfo["base_registry"] = config.RegistryUrl + "/"
	}
	color.Green.Println("如选择错误，请按 Ctrl+C 终止程序，然后重新执行向导")
	challengeInfo["type"] = util.SelectOne("请选择您要创建的题目类型", global.ChallengeType)
	challengeInfo["base_image_name"] = challengeInfo["type"]
	switch challengeInfo["type"] {
	case "web":
		challengeInfo = WizardWeb(challengeInfo)
	case "pwn":
		challengeInfo = WizardPwn(challengeInfo)
	case "misc":
		challengeInfo = WizardSocket(challengeInfo)
	}
	color.Cyan.Println("题目名称应当全为小写，如为中文名称则使用拼音，格式如下")
	color.Cyan.Println("challenge_年份_所属比赛简写_分类_题目名称")
	fmt.Println("")
	color.Cyan.Println("例1 2021年 N1CTF Web类 babysqli，则为 challenge_2021_n1ctf_web_baysqli")
	color.Cyan.Println("例2 2019年 SCTF Pwn类 babyheap，则为 challenge_2019_sctf_pwn_bayheap")
	// 不断获取输入直到有内容
	for {
		challengeInfo["challenge_name"] = util.InputLine("请输入您要创建的题目镜像名称")
		if len(challengeInfo["challenge_name"]) != 0 {
			break
		}
		color.Red.Println("你未输入题目名称，请重新输入")
	}
	// 创建
	Generate(challengeInfo)
}

func WizardWeb(challengeInfo map[string]string) map[string]string {
	// 判断语言
	challengeInfo["language"] = util.SelectOne("请选择您要使用的语言", global.Language)
	// 判断语言版本
	if challengeInfo["language"] != "html" {
		languageVersion := []string{}
		switch challengeInfo["language"] {
		case "php":
			languageVersion = global.PHPVersion
		case "python":
			languageVersion = global.PythonVersion
		case "nodejs":
			languageVersion = global.NodeJSVersion
		case "java":
			languageVersion = global.JavaVersion
		case "ruby":
			languageVersion = global.RubyVersion
		}
		challengeInfo["language_version"] = util.SelectArray("请选择您要使用的版本", languageVersion)
	}
	// 判断Web服务器
	switch challengeInfo["language"] {
	case "php", "html":
		challengeInfo["webserver"] = util.SelectOne("请选择您要使用的Web服务器", global.PHPWebServer)
	case "java":
		challengeInfo["webserver"] = util.SelectOne("请选择您要使用的Web服务器", global.JavaServer)
	case "python":
		challengeInfo["webserver"] = util.SelectOne("请选择您要使用的托管方式", global.PythonWebServer)
	}
	// 判断数据库
	if challengeInfo["language"] != "html" {
		challengeInfo["db"] = util.SelectOne("是否需要数据库", global.DBType)
	}
	// 拼接镜像名称
	baseImageName := ""
	switch challengeInfo["language"] {
	case "php", "python", "java", "html":
		baseImageName += "_" + challengeInfo["webserver"]
	case "nodejs", "ruby":
		// 不需要webserver
	}
	if challengeInfo["db"] != "" {
		baseImageName += "_" + challengeInfo["db"]
	}
	if challengeInfo["language"] != "html" {
		baseImageName += "_" + challengeInfo["language"]
		baseImageName += "_" + challengeInfo["language_version"]
	}
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}

func WizardPwn(challengeInfo map[string]string) map[string]string {
	challengeInfo["pwn_server"] = util.SelectOne("请选择您期望的启动方式", global.PwnServer)
	challengeInfo["pwn_arch"] = util.SelectOne("请选择您要使用的架构", global.PwnArch)
	// 拼接镜像名称
	baseImageName := ""
	switch challengeInfo["pwn_arch"] {
	case "":
		baseImageName += ""
	case "kernel", "qemu":
		baseImageName += "_" + challengeInfo["pwn_arch"]
	}
	baseImageName += "_" + challengeInfo["pwn_server"]
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}

func WizardSocket(challengeInfo map[string]string) map[string]string {
	selectLanguage := global.Language
	delete(selectLanguage, "HTML")
	delete(selectLanguage, "PHP")
	delete(selectLanguage, "Ruby")
	challengeInfo["language"] = util.SelectOne("请选择您要使用的语言", selectLanguage)
	challengeInfo["db"] = util.SelectOne("是否需要数据库", global.DBType)
	// 判断语言版本
	languageVersion := []string{}
	switch challengeInfo["language"] {
	case "python":
		languageVersion = global.PythonVersion
	case "nodejs":
		languageVersion = global.NodeJSVersion
	case "java":
		languageVersion = global.JavaVersion
	}
	challengeInfo["language_version"] = util.SelectArray("请选择您要使用的版本", languageVersion)
	// 拼接镜像名称
	baseImageName := ""
	if challengeInfo["db"] != "" {
		baseImageName += "_" + challengeInfo["db"]
	}
	baseImageName += "_" + challengeInfo["language"]
	baseImageName += "_" + challengeInfo["language_version"]
	// version := util.SelectOne("请选择您要使用的版本", global.SelectVersion[language])
	// hasDB = util.SelectOne("是否需要数据库", []string{"无", "MySQL", "MongoDB"})
	// baseImageName += "_misc_socat"
	// if hasDB != "无" {
	// 	baseImageName += "_" + strings.ToLower(hasDB)
	// 	hasDB = strings.ToLower(hasDB)
	// }
	// baseImageName += strings.ToLower(version)
	challengeInfo["base_image_name"] += baseImageName
	return challengeInfo
}
