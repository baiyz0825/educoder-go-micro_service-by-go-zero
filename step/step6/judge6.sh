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
touch /tmp/judge6-run.log
# 删除数据库user
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_user;'  > /dev/null 2>&1
# 执行用户sql语句
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"    < /data/workspace/myshixun/step/step6/stu/user.sql
{
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_user;describe user;";
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_user;describe third;";
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_user;describe third_data;";
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e "use school_user;describe major;";
} >>  /tmp/judge6-run.log

# 删除数据库
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}"   -e 'drop database if exists school_user;' > /dev/null 2>&1
# 输出结果
cat /tmp/judge6-run.log
rm -f /tmp/judge6-run.log
