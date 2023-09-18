package cmdutil

import (
	"cg/pkg/global"
	"cg/pkg/tpl"
	"cg/pkg/util"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func GetServicePort(challengeType string) string {
	switch challengeType {
	case "web":
		return "80"
	case "pwn":
		return "10000"
	case "misc":
		return "10000"
	case "web_access":
		return "10800"
	default:
		return "9999"
	}
}

/**
 * @Description: 生成 Dockerfile
 * @param challengeInfo 题目信息
 */
func GenerateDockerFile(challengeInfo map[string]string) {
	baseImageName := challengeInfo["base_registry"] + challengeInfo["base_image_name"]
	// Dockerfile
	dockerfile := "FROM " + baseImageName + "\n"
	dockerfile += "MAINTAINER CTFHub Team\n"
	dockerfile += "\n"
	// 根据类型判断题目内容放置位置
	switch challengeInfo["type"] {
	case "web":
		switch challengeInfo["language"] {
		case "php", "html":
			dockerfile += "COPY ./src/ /var/www/html/\n"
		default:
			dockerfile += "COPY ./src/ /app/\n"
		}
	case "pwn":
		dockerfile += "COPY ./src/pwn /pwn\n"
	case "misc":
		dockerfile += "COPY ./src/ /app/\n"
	}
	dockerfile += "\n"
	dockerfile += "EXPOSE " + GetServicePort(challengeInfo["type"]) + "\n"
	util.WriteFile("enviroment/Dockerfile", dockerfile, 0644)
}

/**
 * @Description: 生成 docker-compose.yml
 * @param challengeInfo 题目信息
 */
func GenerateDockerCompose(challengeInfo map[string]string) {
	servicePort := ""
	accessPort := ""
	servicePort = GetServicePort(challengeInfo["type"])
	if challengeInfo["type"] == "web" {
		accessPort = GetServicePort("web_access")
	} else {
		accessPort = servicePort
	}
	dockerCompose := global.DockerCompsoe{}
	_ = yaml.Unmarshal(tpl.DockerCompose, &dockerCompose)
	// 修改内容
	dockerCompose.Version = 3
	dockerCompose.Services.Challenge.Image = challengeInfo["challenge_name"]
	dockerCompose.Services.Challenge.Ports = []string{accessPort + ":" + servicePort}
	dockerCompose.Services.Challenge.Enviroment = []string{
		"FLAG=ctfhub{test_flag}",
		"DOMAIN=test.sandbox.ctfhub.com",
	}
	// 写入文件
	writeData, _ := yaml.Marshal(&dockerCompose)
	util.WriteFile("enviroment/docker-compose.yml", string(writeData), 0644)
}

/**
 * @Description: 生成 meta.yml
 * @param challengeInfo 题目信息
 */
func GenerateMeta(challengeInfo map[string]string) {
	// 读取配置
	config := global.Config{}
	UserHomeDir, _ := os.UserHomeDir()
	data, _ := util.ReadFileByte(UserHomeDir + "/.config/cg/config.yaml")
	_ = yaml.Unmarshal(data, &config)
	// 载入模板
	meta := global.Meta{}
	_ = yaml.Unmarshal(tpl.Meta, &meta)

	// 修改内容
	meta.Author.Name = config.Author
	meta.Author.Contact = config.Contact
	meta.Task.Name = challengeInfo["challenge_name"]
	meta.Task.Type = challengeInfo["type"]
	meta.Task.Level = challengeInfo["level"]
	// 写入文件
	writeData, _ := yaml.Marshal(&meta)
	util.WriteFile("meta.yml", string(writeData), 0644)
}

/**
 * @Description: 生成 flag.sh
 * @param challengeInfo 题目信息
 */
func GenerateFlag(challengeInfo map[string]string) {
	if challengeInfo["need_flag"] == "是" {
		util.WriteFile("enviroment/files/flag.sh", string(tpl.Flag), 0755)
	}
}

/**
 * @Description: 生成 start.sh
 * @param challengeInfo 题目信息
 */
func GenerateStart(challengeInfo map[string]string) {
	if challengeInfo["need_start"] == "是" {
		util.WriteFile("enviroment/files/start.sh", string(tpl.Start), 0755)
	}
}

/**
 * @Description: 生成 db.sql/db.json
 * @param challengeInfo 题目信息
 */
func GenerateDB(challengeInfo map[string]string) {
	switch challengeInfo["db_type"] {
	case "mysql":
		util.WriteFile("enviroment/files/db.sql", string(tpl.DB_SQL), 0644)
	case "mongodb":
		util.WriteFile("enviroment/files/db.json", string(tpl.DB_JSON), 0644)
	}
}

/**
 * @Description: 生成 README.md
 * @param challengeInfo 题目信息
 */
func GenerateReadme(challengeInfo map[string]string) {
	// 读取meta
	meta := global.Meta{}
	_ = yaml.Unmarshal(tpl.Meta, &meta)
	readme := string(tpl.Readme)
	readme = strings.Replace(readme, "CHALLENGE_NAME", meta.Challenge.Name, -1)
	readme = strings.Replace(readme, "CHALLENGE_REFER", meta.Challenge.Refer, -1)
	readme = strings.Replace(readme, "AUTHOR", meta.Author.Name, -1)
	readme = strings.Replace(readme, "EMAIL", meta.Author.Contact, -1)
	readme = strings.Replace(readme, "TASK_NAME", meta.Task.Name, -1)
	readme = strings.Replace(readme, "TASK_TYPE", meta.Task.Type, -1)
	readme = strings.Replace(readme, "TASK_LEVEL", meta.Task.Level, -1)
	if meta.Task.Flag == "" {
		readme = strings.Replace(readme, "TASK_FLAG", "动态 flag", -1)
	} else {
		readme = strings.Replace(readme, "TASK_FLAG", meta.Task.Flag, -1)
	}

	// 写入文件
	util.WriteFile("README.md", string(readme), 0644)

}
