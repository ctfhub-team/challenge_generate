package global

// 镜像源配置
var RegistryNameSpace = "ctfhub"
var RegistryName = []string{"AliYun", "CTFHub", "DockerHub"}
var Registry = map[string]string{
	"AliYun":    "https://registry.cn-hangzhou.aliyuncs.com",
	"CTFHub":    "",
	"DockerHub": "https://hub.docker.com",
}

// 题目类型
var ChallengeType = []string{"Web", "Pwn", "其他"}

// 语言类型
var Language = []string{"PHP", "HTML", "Python", "NodeJS", "Java", "Ruby"}

// PHP Web服务器
var PhpWebServer = []string{"Apache HTTPd", "Nginx"}

// Python Web服务器
var PythonWebServer = []string{"gunicron", "supervisor"}

// Java Web服务器
var JavaServer = []string{"Tomcat"}

// 语言版本
var SelectVersion = map[string][]string{
	"PHP":    {"5.6", "7.4"},
	"Python": {"2.7", "3.8"},
	"NodeJS": {"14.15"},
	"Java":   {"11"},
	"Ruby":   {"2.7"},
}

// Pwn架构
var PwnArch = []string{"x86/x64", "arm/arm64/mips/mips64"}

// Pwn启动方式
var PwnServer = []string{"socat", "xinetd"}

var FileTemplate = map[string]string{
	"flag.sh":  "#!/bin/bash\nflag",
	"start.sh": "#!/bin/bash\nstart",
	"db.sql":   "db",
}
