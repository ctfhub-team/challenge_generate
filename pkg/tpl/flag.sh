#!/bin/bash
<<PROMPT
请根据需要自行修改此脚本，确保flag能够放置在正确位置
部分Web题目需要正确的域名访问，请在此脚本中配置正确的域名

FLAG 平台传入的正确的flag
DOMAIN 平台传入的当前环境的域名

以下为已经内置的函数

将flag写入到文件系统中
# 默认写入/flag
write_flag_in_fs

# 写入至web目录 /var/www/html/flag.txt
write_flag_in_fs /var/www/html/flag.txt

-----------------

将flag写入到数据库
# 默认写入web库中flag表的flag字段
write_flag_in_db

# 指定库为sqli, 表为user, 列为flag写入
write_flag_in_db sqli user flag
PROMPT

write_flag_in_fs() {
    # 将flag写入到文件系统中
    if [ -z "$1" ]; then
        flag_path="/flag"
    else
        flag_path="$1"
    fi
    echo ${FLAG} > ${flag_path}
}

write_flag_in_db() {
    local db_name="${1:-web}"
    local db_table="${2:-flag}"
    local db_column="${3:-flag}"
    echo mysql -uroot -proot -e "update ${db_name}.${db_table} set ${db_column}='${FLAG}';"
}

export FLAG=not_flag
FLAG=not_flag