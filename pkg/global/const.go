package global

// 镜像源配置
var RegistryNameSpace = "ctfhub_base"
var Registry = map[string]string{
	"AliYun":    "https://registry.cn-hangzhou.aliyuncs.com/",
	"CTFHub":    "",
	"DockerHub": "https://hub.docker.com/",
}

// 题目类型
var ChallengeType = map[string]string{
	"Web": "web",
	"Pwn": "pwn",
	// "Socket": "misc",
}

// 语言类型
var Language = map[string]string{
	"PHP":  "php",
	"HTML": "html",
	// "Python": "python",
	// "NodeJS": "nodejs",
	// "Java":   "java",
	// "Ruby":   "ruby",
}

var PHPVersion = []string{
	"5.6", "7.4", "8.0",
}
var PythonVersion = []string{
	"2.7", "3.6",
}
var NodeJSVersion = []string{
	"12", "14", "16", "18",
}
var JavaVersion = []string{
	"8", "11", "15",
}
var RubyVersion = []string{
	"2.5", "2.6", "2.7",
}

var DBType = map[string]string{
	"无/SQLite": "",
	"MySQL":    "mysql",
	// "MongoDB":  "mongodb",
}

// PHP Web服务器
var PHPWebServer = map[string]string{
	"Apache HTTPd": "httpd",
	"Nginx":        "nginx",
}

// Python Web服务器
var PythonWebServer = map[string]string{
	"gunicron":   "gunicron",
	"supervisor": "supervisor",
}

// Java Web服务器
var JavaServer = map[string]string{
	"Tomcat": "tomcat",
}

// Pwn架构
var PwnArch = map[string]string{
	"x86/x64 Binary": "binary",
	// "x86/x64 Kernel":        "kernel",
	// "arm/arm64/mips/mips64": "qemu",
}

// Pwn启动方式
var PwnServer = map[string]string{
	"socat":   "socat",
	"xinet.d": "xinetd",
}

var FileTemplate = map[string]string{
	"flag.sh":  "#!/bin/bash\nflag",
	"start.sh": "#!/bin/bash\nstart",
	"db.sql":   "db",
}

// 难度
var Level = []string{"签到", "简单", "中等", "困难"}
