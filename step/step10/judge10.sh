#!/bin/bash
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
# 检测mysql链接
while true; do
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "show databases;" >/dev/null 2>&1
  if [ $? -eq 0 ]; then
    break
  else
    sleep 1
  fi
done
cd /data/workspace/myshixun/step/step10/stu || exit
# 创建日志文件
touch /tmp/judge10-run.log
# 删除数据库user中所有数据
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e 'drop database user;' >/dev/null 2>&1
# 创建user
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step10/stu/user.sql >/dev/null 2>&1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e 'INSERT INTO user.user (user_name, passwd) VALUES ("bigwhite", "abcd12345h");' >/dev/null 2>&1
# 编译程序，获取输出结果
go run /data/workspace/myshixun/step/step10/stu/main.go >/tmp/judge10-run.log
sleep 10
cat /tmp/judge10-run.log
# 预期结果  user name:bigwhite,passwd:abcd1234
