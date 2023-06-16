#!/bin/bash
source /etc/profile
# 清理环境
cleanEnv() {
  rm /tmp/judge15-res.log >/dev/null 2>&1
  rm /tmp/judge15-api-run.log >/dev/null 2>&1
  rm /tmp/judge15-rpc-run.log >/dev/null 2>&1
}
cleanEnv
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
RPC_URL="localhost:8013"
# 检测mysql链接
while true; do
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "show databases;" >/dev/null 2>&1
  if [ $? -eq 0 ]; then
    break
  else
    sleep 1
  fi
done

# 初始化数据库
if mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "SHOW DATABASES LIKE 'school_user';" | grep -q school_user; then
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_user;" >/dev/null 2>&1
fi
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step15/judge_data/school_user.sql >/dev/null 2>&1
echo "开始测评"
# 创建日志文件
touch /tmp/judge15-rpc-run.log
touch /tmp/judge15-api-run.log
touch /tmp/judge15-res.log

# 编译启动服务
cd /data/workspace/myshixun/sources/apps/user/cmd/rpc || exit
echo "编译代码 user rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/rpc/userrpc.go' >/tmp/judge15-rpc-run.log 2>&1 &
# 等待服务启动

# 编译启动服务
cd /data/workspace/myshixun/sources/apps/user/cmd/api || exit
echo "编译代码 user api"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/api/user.go' >/tmp/judge15-api-run.log 2>&1 &
# 等待服务启动

testing() {
  if timeout 240 tail -f /tmp/judge15-res.log | while true; do
    if grep -q "Start" /tmp/judge15-rpc-run.log && grep -q "Start" /tmp/judge15-api-run.log; then

      # 检查验证码接口 /user/v1/captcha
      echo "检查验证码接口 /user/v1/captcha"
      response=$(curl -s http://127.0.0.1:8003/user/v1/captcha)
      #echo "$response"
      code=$(echo "$response" | sed -n 's/.*"code":\s*\([0-9]\+\).*/\1/p')
      msg=$(echo "$response" | sed -n 's/.*"msg":\s*"\([^"]*\)".*/\1/p')
      captchaB64=$(echo "$response" | sed -n 's/.*"captchaB64":\s*"\([^"]*\)".*/\1/p')
      captchaId=$(echo "$response" | sed -n 's/.*"captchaId":\s*"\([^"]*\)".*/\1/p')
      #echo "$code"
      #echo "$msg"
      #echo "$captchaId"
      #echo "$captchaB64"
      if [ "$code" == 200 ]; then
        echo "code is ok" >>/tmp/judge15-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "message is ok" >>/tmp/judge15-res.log
      fi
      if [ -n "$captchaB64" ]; then
        echo "captchaB64 is ok" >>/tmp/judge15-res.log
      fi
      if [ -n "$captchaId" ]; then
        echo "captchaId is ok" >>/tmp/judge15-res.log
      fi
      # 检查用户注册接口
      echo "检查用户注册接口 /user/v1/register"
      registerResp=$(curl -s --request POST \
        --url http://127.0.0.1:8003/user/v1/register \
        --header 'Accept: application/json' \
        --header 'Content-Type: multipart/form-data' \
        --header 'content-type: multipart/form-data' \
        --form username=test-reg \
        --form password=testAbcdTest111 \
        --form phone=18196984099 \
        --form captchaId=YvqOyYYm1qjfgOkhrxvG \
        --form captcha=46068)
      code=$(echo $registerResp | sed 's/.*"code":\([[:digit:]]\{1,\}\),.*$/\1/')
      msg=$(echo $registerResp | sed 's/.*"msg":"\([^"]\{1,\}\)",.*$/\1/')
      accessToken=$(echo $registerResp | sed 's/.*"accessToken":"\([^"]\{1,\}\)".*$/\1/')
      if [ "$code" == 200 ]; then
        echo "reg code is ok" >>/tmp/judge15-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "reg message is ok" >>/tmp/judge15-res.log
      fi
      if [ -n "$accessToken" ]; then
        echo "reg captchaId is ok" >>/tmp/judge15-res.log
      fi
      sleep 2

      # 获取用户数据
      echo "获取用户数据 /user/v1/info"
      bearer_token="Bearer ${accessToken}" # 添加 Bearer 前缀
      userInfo=$(curl -s --request GET \
        --url http://127.0.0.1:8003/user/v1/info \
        --header 'Accept: application/json' \
        --header "Authorization: ${bearer_token}" \
        --header 'Content-Type: application/json')
      #echo "$userInfo" >> /tmp/judge15-res.log
      code=$(echo $userInfo | sed 's/.*"code":\([0-9]*\).*/\1/')
      msg=$(echo $userInfo | sed 's/.*"msg":"\([^"]*\)".*/\1/')
      phone=$(echo $userInfo | sed 's/.*"phone":"\([^"]*\)".*/\1/')
      if [ "$code" == 200 ]; then
        echo "reg code is ok" >>/tmp/judge15-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "reg message is ok" >>/tmp/judge15-res.log
      fi
      phoneStr=$(printf '"%d"' "$phone")
      if [ "$phoneStr" == "18196984099" ]; then
        echo "reg phone is ok" >>/tmp/judge15-res.log
      fi
      echo "reg code=$code,msg=$msg, phone=$phone" >>/tmp/judge15-res.log

      # kill user-api
      kill $(lsof -t -i:8003) >/dev/null 2>&1
      # kill user-rpc
      kill $(lsof -t -i:8013) >/dev/null 2>&1
      user_res="/tmp/judge15-res.log"
      ans_res="/data/workspace/myshixun/step/step15/judge_ans/ans.log"
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $user_res
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $ans_res
      differAns="$(diff $user_res $ans_res)"
      if [ -z "$differAns" ]; then
        echo "Yes"
        echo "评测结束"
        cleanEnv
        exit
      else
        echo "评测失败"
        echo "测试输出为:"
        cat /tmp/judge15-res.log
        echo "user-api程序输出为:"
        cat /tmp/judge15-api-run.log
        echo "user-rpc程序输出为:"
        cat /tmp/judge15-rpc-run.log
        cleanEnv
        exit
      fi
    fi
  done; then
    cleanEnv
  else
    echo "测评超时"
    cleanEnv
  fi
}

# 启动后台任务执行函数
testing &
# 记录后台任务的pid
job_pid=$!
# 等待后台任务执行完毕
wait $job_pid
# 评测正确答案
#mysql: [Warning] Using a password on the command line interface can be insecure.
#开始测评
#编译代码 user rpc
#编译代码 user api
#检查验证码接口 /user/v1/captcha
#检查用户注册接口 /user/v1/register
#获取用户数据 /user/v1/info
#Yes
#评测结束
