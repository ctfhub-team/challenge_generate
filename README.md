# challenge_generate

题目环境生成器

基于向导或是配置文件进行题目整体框架创建的工具

```bash
Challenge Generate (cg) 是用于创建CTF题目环境模板的自助式生成工具.

https://www.ctfhub.com/            Developed by: L1n3@CTFHub Team

Usage:
  cg [command]

Available Commands:
  check        检查当前题目目录及内容是否符合规范
  completion   Generate the autocompletion script for the specified shell
  config       设置 cg 的配置文件
  docker       docker相关操作
  help         Help about any command
  new          创建新的题目模板
  version      输出 cg 的版本和更新时间

Flags:
  -h, --help   help for cg

Use "cg [command] --help" for more information about a command.
```

## 创建

使用`new`子命令进行创建，分为向导模式和文件模式
```bash
n 创建新的题目模板

Usage:
  cg new [flags]
  cg new [command]

Aliases:
  new, n

Available Commands:
  file         从预定义文件创建
  wizard       使用向导创建

Flags:
  -h, --help   help for new

Use "cg new [command] --help" for more information about a command.
```

## 测试

测试需要将当前目录切换至含有`docker-compose.yml`文件的目录，一般来说即为题目目录中的`environment` 目录，之后执行`docker`子命令按需处理即可

```bash
d 测试已完成的镜像

Usage:
  cg docker [flags]
  cg docker [command]

Aliases:
  docker, d

Available Commands:
  auto         依次执行 Stop -> Build -> Run -> Bash
  bash         执行bash进入容器
  build        构建镜像
  log          查看容器日志
  run          运行镜像
  save         导出容器tar包
  stop         停止镜像

Flags:
  -h, --help   help for docker

Use "cg docker [command] --help" for more information about a command.
```

## 环境目录结构

```
environment
    src/
    files/
        flag.sh
        db.sql
        start.sh
    Dockerfile
    docker-compose.yml
writeup/
    images/
    readme.md
    writeup.pdf
    exp.py
    requirement.txt
attachment.zip
meta.yml
```