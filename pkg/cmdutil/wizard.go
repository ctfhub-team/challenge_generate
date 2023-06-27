package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/util"
	"fmt"
	"os"
	"strings"
)

/**
 * @Description: 生成题目模板
 * @param baseImageName 题目基础镜像名称
 * @param challengeType 题目类型
 * @param challengeName 题目名称
 * @param language 题目语言
 * @param hasDB 是否需要数据库
 */
func Generate(baseImageName string, challengeType string, challengeName string, language string, hasDB string) {
	// flag位置确认
	needFlag := util.SelectOne("是否需要单独处理flag位置", []string{"否", "是"})

	// 二次确认
	confirm := util.SelectOne("确认创建题目 "+challengeName+" ?", []string{"确认", "取消"})
	if confirm == "取消" {
		os.Exit(0)
	}

	// 服务端口
	servicePort := 0
	switch challengeType {
	case "Web":
		servicePort = 80
	case "Pwn":
		servicePort = 10000
	case "其他":
		servicePort = 10000
	}

	// 创建题目模板目录
	os.Mkdir(challengeName, 0755)
	os.Chdir(challengeName)

	// 创建文件夹
	dirTree := []string{
		"enviroment/src/",
		"enviroment/files/",
		"writeup/",
	}
	for _, path := range dirTree {
		os.MkdirAll(path, os.ModePerm)
	}
	// flag.sh
	if needFlag == "是" {
		flagFile, _ := os.OpenFile("enviroment/files/flag.sh", os.O_CREATE|os.O_WRONLY, 0755)
		flagFile.WriteString(global.FileTemplate["flag.sh"])
	}

	// db.sql
	if hasDB != "无" {
		dbFile, _ := os.OpenFile("enviroment/files/db.sql", os.O_CREATE|os.O_WRONLY, 0755)
		dbFile.WriteString(global.FileTemplate["db.sql"])
	}
	// start.sh
	startFile, _ := os.OpenFile("enviroment/files/start.sh", os.O_CREATE|os.O_WRONLY, 0755)
	startFile.WriteString(global.FileTemplate["start.sh"])

	// docker-compose.yml
	dockerCompsoe := "version: \"3\"\n"
	dockerCompsoe += "services:\n"
	dockerCompsoe += "  challenge:\n"
	dockerCompsoe += "    build: .\n"
	dockerCompsoe += "    image: " + strings.ToLower(challengeName) + "\n"
	dockerCompsoe += "    ports:\n"
	if challengeType == "Web" {
		dockerCompsoe += "      -  10800:80\n"
	} else {
		dockerCompsoe += "      -  10000:10000\n"
	}
	dockerCompsoe += "    environment:\n"
	dockerCompsoe += "      - FLAG=ctfhub{test_flag}\n"
	dockerCompsoe += "      - DOMAIN=test.sandbox.ctfhub.com"
	dockerComposeFile, _ := os.OpenFile("enviroment/docker-compose.yml", os.O_CREATE|os.O_WRONLY, 0644)
	dockerComposeFile.WriteString(dockerCompsoe)

	// Dockerfile
	dockerfile := "FROM " + baseImageName + "\n"
	dockerfile += "MAINTAINER CTFHub Team\n"
	dockerfile += "\n"
	// 根据类型判断题目内容放置位置
	switch challengeType {
	case "Web":
		switch language {
		case "PHP", "HTML":
			dockerfile += "COPY ./src/ /var/www/html/\n"
		default:
			dockerfile += "COPY ./src/ /app/\n"
		}
	case "Pwn":
		dockerfile += "COPY ./src/pwn /pwn/pwn\n"
	case "其他":
		dockerfile += "COPY ./src/ /app/\n"
	}

	if needFlag == "是" {
		dockerfile += "COPY ./files/flag.sh /flag.sh\n"
	}
	if hasDB != "无" {
		switch hasDB {
		case "mysql":
			dockerfile += "COPY ./files/db.sql /db.sql\n"
		case "mongodb":
			dockerfile += "COPY ./files/db.json /db.json\n"
		}
	}
	dockerfile += "\n"
	dockerfile += "EXPOSE " + fmt.Sprintf("%d", servicePort) + "\n"
	dockerFile, _ := os.OpenFile("enviroment/Dockerfile", os.O_CREATE|os.O_WRONLY, 0755)
	dockerFile.WriteString(dockerfile)

	// TODO meta.yml生成

	// 输出成功提示
	fmt.Println("")
	fmt.Println(challengeName + " 创建成功，请按如下步骤依次操作：")
	fmt.Println("1. 初始化Git仓库")
	fmt.Println("2. 编辑 " + challengeName + "/meta.yml 文件修改题目信息")
	fmt.Println("3. 进入 " + challengeName + "/enviroment 目录进行测试")
	fmt.Println("4. 进入 " + challengeName + "/writeup 目录补全WP及Exp")
	fmt.Println("5. 一切完成后推送至Git仓库")
}

func Wizard() {
	hasDB := "无"
	language := ""
	registry := util.SelectOne("请选择您要使用的镜像源", global.RegistryName)
	challengeType := util.SelectOne("请选择您要创建的题目类型", global.ChallengeType)
	baseImageName := global.Registry[registry] + "/" + global.RegistryNameSpace + "/"
	baseImageName += "base"
	switch challengeType {
	case "Web":
		// 判断语言
		language = util.SelectOne("请选择您要使用的语言", global.Language)
		// 判断语言版本
		version := ""
		if language == "HTML" {
			version = ""
		} else {
			version = util.SelectOne("请选择您要使用的版本", global.SelectVersion[language])
		}
		// 判断Web服务器
		webServer := ""
		switch language {
		case "PHP", "HTML":
			webServer = util.SelectOne("请选择您要使用的Web服务器", global.PhpWebServer)
			if webServer == "Apache HTTPd" {
				webServer = "httpd"
			}
		case "Java":
			webServer = util.SelectOne("请选择您要使用的Web服务器", global.JavaServer)
		case "Python":
			webServer = util.SelectOne("请选择您要使用的托管方式", global.PythonWebServer)
		default:
			webServer = ""
		}
		// 判断数据库
		if language != "HTML" {
			hasDB = util.SelectOne("是否需要数据库", []string{"SQLite/无", "MySQL"})
			if hasDB == "MySQL" {
				hasDB = "mysql"
			} else {
				hasDB = "无"
			}
		}
		// 拼接镜像名称
		baseImageName += "_web"
		baseImageName += "_" + strings.ToLower(webServer)
		if language != "HTML" {
			if hasDB != "无" {
				baseImageName += "_" + strings.ToLower(hasDB)
			}
			baseImageName += "_" + strings.ToLower(language)
			baseImageName += "_" + version
		}
	case "Pwn":
		server := util.SelectOne("请选择您期望的启动方式", global.PwnServer)
		arch := util.SelectOne("请选择您要使用的架构", global.PwnArch)
		baseImageName += "_pwn"
		baseImageName += strings.ToLower(server)
		if arch == "arm/arm64/mips/mips64" {
			baseImageName += "_qemu"
		}
	case "其他":
		language = util.SelectOne("请选择您要使用的语言", global.Language)
		if language == "PHP" {
			fmt.Println("暂不支持PHP")
			os.Exit(0)
		}
		version := util.SelectOne("请选择您要使用的版本", global.SelectVersion[language])
		hasDB = util.SelectOne("是否需要数据库", []string{"无", "MySQL", "MongoDB"})
		baseImageName += "_misc_socat"
		if hasDB != "无" {
			baseImageName += "_" + strings.ToLower(hasDB)
			hasDB = strings.ToLower(hasDB)
		}
		baseImageName += strings.ToLower(version)

	}
	challengeName := util.InputLine("请输入您要创建的题目镜像名称")
	// 创建
	Generate(baseImageName, challengeType, strings.TrimSpace(challengeName), language, hasDB)
}
