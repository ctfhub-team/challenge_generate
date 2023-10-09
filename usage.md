## 详细使用说明

https://www.wolai.com/ctfhub/3DvnJJtPbHyyDtVkDaW1yz

## 要求

所有环境基于amd64构建

## 创建题目

请按照如下流程进行创建

0. 配置制作者信息

制作者信息仅需配置一次，如配置错误请使用`cg config clean`命令清除后重新配置

```bash
cg config set
```

1. 使用向导生成题目模板

使用`new`子命令创建
```bash
cg new wizard
```

如果对flag放置位置有特殊要求，请在向导中`单独处理flag位置`选项选择`是`，创建完成后修改`environment/files/flag.sh`文件

如对环境的服务启动有特殊要求，请在向导中`单独处理部分服务启动`选项选择`是`,创建完成后修改`environment/files/start.sh`文件

2. 参考示例完成镜像

请确保所使用的基础镜像为最新，可在build之前自行pull一次对应的基础镜像

生成后如无特殊需求，请尽可能不要改动dockerfile

示例链接: https://github.com/ctfhub-team/base_image/

3. 测试镜像

使用`docker`子命令进行环境测试
```bash
cg docker auto
```

4. 修改元信息

编辑题目目录中的`meta.yaml`，具体含义请参考`meta说明`

5. 打包

使用zip打包整个题目文件夹

```bash
zip -r xxx.zip xxx
```

6. 上传

请将打包好的题目文件夹传给相关负责人

## meta说明

```yaml
author:
  # 制作者ID，由cg自动生成
  name: l1n3
  # 制作者邮箱，由cg自动生成
  contact: yw9381@163.com
task:
  # 题目镜像名称，由cg自动生成
  name: challenge_web_2022_hitcon_rce
  # 题目类型，由cg自动生成
  type: Web
  # 题目描述
  description: aasasdsadas
  # 题目难度，由cg自动生成
  level: 签到
  # 题目flag
  # 如是静态flag在此处填写具体的flag值
  # 如是动态flag则此处为空字符串
  flag: 
  # 题目提示，如无提示则此处为空数组
  hints:
    - asdasd
    - asdas
    - asdad
challenge:
  # 题目显示名称
  name: rce
  # 题目来源，来源格式: 年份-比赛名称简写-题目类型-题目显示名称
  # 例如 2021年强网杯的Web类的babysqli，则为2021-QWB-Web-baysqli
  # 例如 2019年SCTF的Pwn类的babyheap，则为2019-SCTF-Pwn-bayheap
  refer: 2022-HitCon-Web-rce
  # 题目标签，标签应当尽可能体现题目考点，如无标签则此处为空数组
  tags:
    - web
    - 2022
    - hitcon

```
## 已知问题

1. Python/NodeJS/Ruby/Java 基础环境暂未完成
2. check功能暂未完成，如制作完成后请直接将打包好的文件发送给对接人，由对接人负责审核测试