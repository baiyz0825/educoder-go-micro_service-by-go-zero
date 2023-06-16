#!/bin/bash
source /etc/profile
# 清理环境
cleanEnv() {
  # kill service
  # kill service
  kill $(lsof -t -i:8013) >/dev/null 2>&1 # user-rpc
  kill $(lsof -t -i:8012) >/dev/null 2>&1 # trade-rpc
  kill $(lsof -t -i:8011) >/dev/null 2>&1 # resources-rpc
  kill $(lsof -t -i:8010) >/dev/null 2>&1 # order-rpc
  kill $(lsof -t -i:8000) >/dev/null 2>&1 # order-api
  kill $(lsof -t -i:8003) >/dev/null 2>&1 # user-api
  rm /tmp/judge18-user-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge18-user-api-run.log >/dev/null 2>&1
  rm /tmp/judge18-resources-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge18-trade-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge18-order-rpc-run.log >/dev/null 2>&1
  rm /tmp/judge18-order-api-run.log >/dev/null 2>&1
  rm /tmp/judge18-res.log >/dev/null 2>&1
}
# kill service
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
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_order;" >/dev/null 2>&1

sleep 2
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step18/judge_data/school_user.sql >/dev/null 2>&1
sleep 1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step18/judge_data/school_resources.sql >/dev/null 2>&1
sleep 1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step18/judge_data/school_trade.sql >/dev/null 2>&1
sleep 1
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step18/judge_data/school_order.sql >/dev/null 2>&1
sleep 1

echo "数据库初始化成功"
echo "初始化日志文件"
# 创建日志文件
touch /tmp/judge18-user-rpc-run.log
touch /tmp/judge18-user-api-run.log
touch /tmp/judge18-resources-rpc-run.log
touch /tmp/judge18-trade-rpc-run.log
touch /tmp/judge18-order-rpc-run.log
touch /tmp/judge18-order-api-run.log
touch /tmp/judge18-res.log

# 编译服务
# user-rpc
cd /data/workspace/myshixun/sources/apps/user/cmd/rpc || exit
echo "编译代码 user rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/rpc/userrpc.go' >/tmp/judge18-user-rpc-run.log 2>&1 &

# resources-rpc
cd /data/workspace/myshixun/sources/apps/resources/cmd/rpc || exit
echo "编译代码 resources rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/resources/cmd/rpc/resourcesrpc.go' >/tmp/judge18-resources-rpc-run.log 2>&1 &

# trade-rpc
cd /data/workspace/myshixun/sources/apps/trade/cmd/rpc || exit
echo "编译代码 trade rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/trade/cmd/rpc/traderpc.go' >/tmp/judge18-trade-rpc-run.log 2>&1 &

# order-rpc
cd /data/workspace/myshixun/sources/apps/order/cmd/rpc || exit
echo "编译代码 order rpc"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/order/cmd/rpc/orderrpc.go' >/tmp/judge18-order-rpc-run.log 2>&1 &

# user-api
cd /data/workspace/myshixun/sources/apps/user/cmd/api || exit
echo "编译代码 user api"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/api/user.go' >/tmp/judge18-user-api-run.log 2>&1 &

# order-api
cd /data/workspace/myshixun/sources/apps/order/cmd/api || exit
echo "编译代码 order api"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/order/cmd/api/order.go' >/tmp/judge18-order-api-run.log 2>&1 &

touch /tmp/judge18-user-rpc-run.log
touch /tmp/judge18-user-api-run.log
touch /tmp/judge18-resources-rpc-run.log
touch /tmp/judge18-trade-rpc-run.log
touch /tmp/judge18-order-rpc-run.log
touch /tmp/judge18-order-api-run.log
touch /tmp/judge18-res.log
# 等待服务启动
testing() {
  if timeout 240 tail -f /tmp/judge18-res.log | while true; do
    if grep -q "Start" /tmp/judge18-user-rpc-run.log &&
      grep -q "Start" /tmp/judge18-user-api-run.log &&
      grep -q "Start" /tmp/judge18-resources-rpc-run.log &&
      grep -q "Start" /tmp/judge18-trade-rpc-run.log &&
      grep -q "Start" /tmp/judge18-order-api-run.log &&
      grep -q "Start" /tmp/judge18-order-rpc-run.log; then
      # 1. 登陆获取token
      echo "测试用户登陆接口: /user/v1/login"
      loginData=$(curl -s --request POST \
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
        echo "login code is ok" >>/tmp/judge18-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "login message is ok" >>/tmp/judge18-res.log
      fi

      bearer_token="Bearer ${accessToken}" # 添加 Bearer 前缀

      # 测试下单接口
      echo "测试下单接口: /trade/v1/order/do"
      doOrder=$(curl -s --request POST \
        --url http://127.0.0.1:8000/trade/v1/order/do \
        --header 'Accept: application/json' \
        --header "Authorization: ${bearer_token}" \
        --header 'Content-Type: application/json' \
        --header 'content-type: application/json' \
        --data '{
                      "productId": 1,
                      "payPath": 1
                  }')

      code=$(echo "$doOrder" | sed 's/.*"code":\([0-9]*\).*/\1/')
      msg=$(echo "$doOrder" | sed 's/.*"msg":"\([^"]*\)".*/\1/')
      payPathOrderNum=$(echo $doOrder | sed 's/.*"payPathOrderNum":"\([^"]*\)".*/\1/')
      if [ "$code" == 200 ]; then
        echo "doOrder code is ok" >>/tmp/judge18-res.log
      fi
      if [ "$msg" == "OK" ]; then
        echo "doOrder message is ok" >>/tmp/judge18-res.log
      fi
      #echo "$doOrder"

      # 查询订单详细信息
      echo "测试查询订单详细信息: /user/v1/login"
      # shellcheck disable=SC2027
      #echo "$payPathOrderNum"
      curl_url="http://127.0.0.1:8000/trade/v1/order/info?uuid=${payPathOrderNum}&userId=3"
      #echo "$curl_url"
      orderInfo=$(curl -s --request GET \
        --url "${curl_url}" \
        --header 'Accept: application/json' \
        --header "Authorization: ${bearer_token}" \
        --header 'Content-Type: application/json' \
        --header 'content-type: application/json')
      # 获取 code
      code=$(echo "$orderInfo" | sed 's/.*"code":\([^,}]*\).*/\1/')
      echo "请求状态为:$code" >>/tmp/judge18-res.log
      # 获取 productOwnerName
      productOwnerName=$(echo "$orderInfo" | sed 's/.*"productOwnerName":"\([^"]*\)".*/\1/')
      echo "产品所属于人名称是:$productOwnerName" >>/tmp/judge18-res.log
      # 获取 TestUploadProduct
      product=$(echo "$orderInfo" | sed 's/.*"name":"\([^"]*\)".*/\1/')
      echo "产品名称是:$product" >>/tmp/judge18-res.log
      # 获取 productOwnerId
      productOwnerId=$(echo "$orderInfo" | sed 's/.*"productOwnerId":\([^,}]*\).*/\1/')
      echo "产品id所属用户id是:$productOwnerId" >>/tmp/judge18-res.log

      # 测试完成
      echo "测试完成,开始比对数据"
      user_res="/tmp/judge18-res.log"
      ans_res="/data/workspace/myshixun/step/step18/judge_ans/ans.log"
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
        cat /tmp/judge18-res.log
        echo "user-rpc程序输出为:"
        cat /tmp/judge18-user-rpc-run.log
        echo "user-api程序输出为:"
        cat /tmp/judge18-user-api-run.log
        echo "trade-rpc程序输出为:"
        cat /tmp/judge18-trade-rpc-run.log
        echo "resources-rpc程序输出为:"
        cat /tmp/judge18-resources-rpc-run.log
        echo "order-rpc程序输出为:"
        cat /tmp/judge18-order-rpc-run.log
        echo "order-api程序输出为:"
        cat /tmp/judge18-order-api-run.log
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
# 测试正确结果输出
#开始测评
#初始化数据库配置
#数据库初始化成功
#初始化日志文件
#编译代码 user rpc
#编译代码 resources rpc
#编译代码 trade rpc
#编译代码 order rpc
#编译代码 user api
#编译代码 order api
#测试用户登陆接口: /user/v1/login
#测试下单接口: /trade/v1/order/do
#测试查询订单详细信息: /user/v1/login
#测试完成,开始比对数据
#Yes,测试通过
#评测结束
