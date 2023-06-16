#!/bin/bash
source /etc/profile
# 清理环境
cleanEnv() {
  # kill service
  kill $(lsof -t -i:8002) >/dev/null 2>&1 # trade-api
  kill $(lsof -t -i:8011) >/dev/null 2>&1 # resources-rpc
  kill $(lsof -t -i:8003) >/dev/null 2>&1 # user-api
  kill $(lsof -t -i:8013) >/dev/null 2>&1 # user-rpc
  kill $(lsof -t -i:8012) >/dev/null 2>&1 # trade-rpc
  rm /tmp/judge17-user-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge17-user-api-run.log >/dev/null 2>&1
  rm /tmp/judge17-resources-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge17-trade-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge17-trade-api-run.log >/dev/null 2>&1
  rm /tmp/judge17-res.log >/dev/null 2>&1
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
echo "开始测评"

echo "初始化数据库配置"
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_user;" >/dev/null 2>&1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_resources;" >/dev/null 2>&1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_trade;" >/dev/null 2>&1

sleep 2
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step17/judge_data/school_user.sql >/dev/null 2>&1
sleep 1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step17/judge_data/school_resources.sql >/dev/null 2>&1
sleep 1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step17/judge_data/school_trade.sql >/dev/null 2>&1
sleep 1

echo "数据库初始化成功"
echo "初始化日志文件"
# 创建日志文件
touch /tmp/judge17-user-rpc-run.log
touch /tmp/judge17-user-api-run.log
touch /tmp/judge17-resources-rpc-run.log
touch /tmp/judge17-trade-rpc-run.log
touch /tmp/judge17-trade-api-run.log
touch /tmp/judge17-res.log

cd /data/workspace/myshixun/sources/apps/resources/cmd/rpc || exit
echo "编译代码 resources rpc "
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/resources/cmd/rpc/resourcesrpc.go' >/tmp/judge17-resources-rpc-run.log 2>&1 &

cd /data/workspace/myshixun/sources/apps/trade/cmd/rpc || exit
echo "编译代码 trade rpc "
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/trade/cmd/rpc/traderpc.go' >/tmp/judge17-trade-rpc-run.log 2>&1 &

cd /data/workspace/myshixun/sources/apps/trade/cmd/api || exit
echo "编译代码 trade api "
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/trade/cmd/api/trade.go' >/tmp/judge17-trade-api-run.log 2>&1 &

# 编译用户服务 user rpc
cd /data/workspace/myshixun/sources/apps/user/cmd/rpc || exit
echo "编译代码 user rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/rpc/userrpc.go' >/tmp/judge17-user-rpc-run.log 2>&1 &

# 编译启动服务 user api
cd /data/workspace/myshixun/sources/apps/user/cmd/api || exit
echo "编译代码 user api"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/api/user.go' >/tmp/judge17-user-api-run.log 2>&1 &

# 等待服务启动
testing() {
  if timeout 240 tail -f /tmp/judge17-res.log | while true; do
    if grep -q "Start" /tmp/judge17-user-rpc-run.log &&
      grep -q "Start" /tmp/judge17-user-api-run.log &&
      grep -q "Start" /tmp/judge17-resources-rpc-run.log &&
      grep -q "Start" /tmp/judge17-trade-rpc-run.log &&
      grep -q "Start" /tmp/judge17-trade-api-run.log; then
      # 1. 登陆获取token
      echo "测试用户登陆接口: /user/v1/login"
      loginData=$(curl -s -s --request POST \
        --url http://127.0.0.1:8003/user/v1/login \
        --header 'Accept: application/json' \
        --header 'Content-Type: multipart/form-data' \
        --header 'content-type: multipart/form-data' \
        --form phone=18196984099 \
        --form password=testAbcdTest111 \
        --form captcha=85438 \
        --form captchaId=eVhvg0n0g7bvd3VDO0Tg)

      # 获取code
      code=$(echo "$loginData" | sed -n 's/.*"code":\([0-9]*\).*/\1/p')
      # 获取msg
      msg=$(echo "$loginData" | sed -n 's/.*"msg":"\([^"]*\)".*/\1/p')
      # 获取accessToken
      accessToken=$(echo "$loginData" | sed -n 's/.*"accessToken":"\([^"]*\)".*/\1/p')
      if [ "$code" == 200 ]; then
        echo "login code is ok" >>/tmp/judge17-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "login message is ok" >>/tmp/judge17-res.log
      fi

      bearer_token="Bearer ${accessToken}" # 添加 Bearer 前缀

      # 测试增加商品接口
      echo "测试增加商品接口接口: /trade/v1/product/addone"
      addProductResp=$(curl -s --request POST \
        --url http://127.0.0.1:8002/trade/v1/product/addone \
        --header 'Accept: application/json' \
        --header "Authorization: ${bearer_token}" \
        --header 'Content-Type: application/json' \
        --header 'content-type: application/json' \
        --data '{
                     "name": "TestUploadProduct",
                     "priductBind": 17,
                     "price": 8.91
                 }')
      code=$(echo "$addProductResp" | sed 's/.*"code":\([0-9]*\),.*/\1/g')
      msg=$(echo "$addProductResp" | sed 's/.*"msg":"\([^"]*\)".*/\1/g')
      if [ "$code" == 200 ]; then
        echo "addone code is ok" >>/tmp/judge17-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "addone message is ok" >>/tmp/judge17-res.log
      fi

      # 分页查询商品
      searchProdut=$(curl -s --request POST \
        --url http://127.0.0.1:8002/trade/v1/product/search \
        --header 'Accept: application/json' \
        --header "Authorization: ${bearer_token} " \
        --header 'Content-Type: application/json' \
        --data '{
                   "page": 1,
                   "limit": 20,
                   "bottonPrice": 0,
                   "topPrice": 10,
                   "desc": true
               }')
      name=$(echo "$searchProdut" | sed 's/.*"name":"\([^"]*\)".*/\1/')
      if [ "$name" == "TestUploadProduct" ]; then
        echo "查询商品数据成功:" >>/tmp/judge17-res.log
        echo "$name" >>/tmp/judge17-res.log
      fi

      # 测试完成
      echo "测试完成,开始比对数据"
      user_res="/tmp/judge17-res.log"
      ans_res="/data/workspace/myshixun/step/step17/judge_ans/ans.log"
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $user_res
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $ans_res
      differAns="$(diff $user_res $ans_res)"
      if [ -z "$differAns" ]; then
        echo "Yes,测试通过"
        echo "评测结束"
        cleanEnv
        exit
      else
        echo "评测失败"
        echo "测试输出为:"
        cat /tmp/judge17-res.log
        echo "user-rpc程序输出为:"
        cat /tmp/judge17-user-rpc-run.log
        echo "user-api程序输出为:"
        cat /tmp/judge17-user-api-run.log
        echo "trade-api程序输出为:"
        cat /tmp/judge17-trade-api-run.log
        echo "trade-rpc程序输出为:"
        cat /tmp/judge17-trade-rpc-run.log
        echo "resources-rpc程序输出为:"
        cat /tmp/judge17-resources-rpc-run.log
        cleanEnv
        exit
      fi
    fi
  done; then
    cleanEnv
  else
    echo "测试超时"
    cleanEnv
  fi
}

# 启动后台任务执行函数
testing &
# 记录后台任务的pid
job_pid=$!
# 等待后台任务执行完毕
wait $job_pid

# 测试成功结果
#开始测评
#初始化数据库配置
#数据库初始化成功
#初始化日志文件
#编译代码 resources rpc
#编译代码 trade rpc
#编译代码 trade api
#编译代码 user rpc
#编译代码 user api
#测试用户登陆接口: /user/v1/login
#测试增加商品接口接口: /trade/v1/product/addone
#测试完成,开始比对数据
#Yes,测试通过
#评测结束

