#!/bin/bash
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
# 检测mysql链接
while true;
do
	mysql  -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"  -e "show databases;" >/dev/null 2>&1
	if [ $? -eq 0 ];then
    	break
	else
    	sleep 1
	fi
done
# 创建日志文件
touch /tmp/judge8-run.log
# 删除数据库user
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_trade;'  > /dev/null 2>&1
# 执行用户sql语句
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"    < /data/workspace/myshixun/step/step8/stu/trade.sql

mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_trade;describe product;" >>  /tmp/judge8-run.log

# 删除数据库
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_trade;' > /dev/null 2>&1
# 输出结果
cat /tmp/judge8-run.log
rm -f /tmp/judge8-run.log
