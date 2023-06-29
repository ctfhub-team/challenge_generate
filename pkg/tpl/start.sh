#!/bin/bash
<<PROMPT
# 请在此脚本中编写要启动的服务，服务应当以后台启动，例如下列例子
redis-server /etc/redis.conf &
PROMPT