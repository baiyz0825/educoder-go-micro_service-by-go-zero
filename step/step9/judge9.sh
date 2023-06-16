#!/bin/bash
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
# 检测mysql链接
while true;
do
	mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"  -e "show databases;" >/dev/null 2>&1
	if [ $? -eq 0 ];then
    	break
	else
    	sleep 1
	fi
done
# 创建日志文件
touch /tmp/judge9-run.log
# 删除数据库user
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_order;'  > /dev/null 2>&1
# 执行用户sql语句
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"    < /data/workspace/myshixun/step/step9/stu/order.sql
{
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_order;describe order;";
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_order;describe user_earn;";
} >>  /tmp/judge9-run.log

# 删除数据库
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_order;' > /dev/null 2>&1
# 输出结果
cat /tmp/judge9-run.log
rm -f /tmp/judge9-run.log
