#!/bin/bash
source /etc/profile
# 清理环境
cleaEnv() {
  # kill order
  kill $(lsof -t -i:8010) >/dev/null 2>&1
  # kill trade
  kill $(lsof -t -i:8012) >/dev/null 2>&1
  rm /tmp/judge14-res.log >/dev/null 2>&1
  rm /tmp/judge14-run-trade.log >/dev/null 2>&1
  rm /tmp/judge14-run-order.log >/dev/null 2>&1
}
cleaEnv
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
RPC_URL="localhost:8012"
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
echo "初始化数据库"
if mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "SHOW DATABASES LIKE 'school_order';" | grep -q school_order; then
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_order;" >/dev/null 2>&1
  sleep 2
fi
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step14/judge_data/school_order.sql >/dev/null 2>&1
echo "开始测评"
touch /tmp/judge14-run-trade.log
touch /tmp/judge14-run-order.log
touch /tmp/judge14-res.log
echo "编译代码"
echo "编译trade-rpc"
cd /data/workspace/myshixun/sources/apps/trade/cmd/rpc || exit
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/trade/cmd/rpc/traderpc.go' >/tmp/judge14-run-trade.log 2>&1 &
echo "编译order-rpc"
cd /data/workspace/myshixun/sources/apps/order/cmd/rpc || exit
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/order/cmd/rpc/orderrpc.go' >/tmp/judge14-run-order.log 2>&1 &
# 等待服务启动

testing() {
  if timeout 240 tail -f /tmp/judge14-res.log | while true; do
    if grep -q "Start" /tmp/judge14-run-trade.log && grep -q "Start" /tmp/judge14-run-order.log; then
      #pb.orderrpc.AddOrder
      #pb.orderrpc.AddUserEarn
      #pb.orderrpc.CheckAilPayStatus
      #pb.orderrpc.DelOrder
      #pb.orderrpc.DelUserEarn
      #pb.orderrpc.DoOrder
      echo -e "评测pb.orderrpc.DoOrder:" >>/tmp/judge14-res.log
      doOrderRes=$(grpcurl -plaintext -d '{"ProductId":6,"UserId":1,"PayPath":1}' localhost:8010 pb.orderrpc.DoOrder)
      hasPayPathOrderNum=$(echo "$doOrderRes" | grep -q 'PayPathOrderNum' && echo 'true' || echo 'false')
      hasStatus=$(echo "$doOrderRes" | grep -q 'Status' && echo 'true' || echo 'false')
      hasPayUrl=$(echo "$doOrderRes" | grep -q 'PayUrl' && echo 'true' || echo 'false')
      if [[ $hasPayPathOrderNum == 'true' && $hasStatus == 'true' && $hasPayUrl == 'true' ]]; then
        echo "创建订单成功" >>/tmp/judge14-res.log
      else
        echo "创建订单失败" >>/tmp/judge14-res.log
        echo "$doOrderRes" >>/tmp/judge14-res.log
      fi
      payPathOrderNum=$(echo "$doOrderRes" | grep -oP '(?<="PayPathOrderNum": ")[^"]*')
      sleep 1
      #pb.orderrpc.GetOrderById
      echo -e "评测pb.orderrpc.GetOrderById:" >>/tmp/judge14-res.log
      getOrderByIdData=$(grpcurl -plaintext -d '{"id":1}' localhost:8010 pb.orderrpc.GetOrderById)
      uuid=$(echo "$getOrderByIdData" | awk -F'"' '/"uuid"/{print $4}')
      payPathOrderNumGetOrderDataNum=$(echo "$getOrderByIdData" | awk -F'"' '/"payPathOrderNum"/{print $4}')
      userId=$(echo "$getOrderByIdData" | awk -F'"' '/"userId"/{print $4}')

      if [ "$payPathOrderNum" = "$payPathOrderNumGetOrderDataNum" ]; then
        echo "订单详情查询成功" >>/tmp/judge14-res.log
      else
        echo "订单详情查询失败" >>/tmp/judge14-res.log
        echo "$getOrderByIdData" >>/tmp/judge14-res.log
      fi
      sleep 1
      #pb.orderrpc.DeleteOrderAliAndDb
      #echo "$uuid"
      #echo "$userId"
      #echo "$payPathOrderNumGetOrderDataNum"
      echo -e "评测pb.orderrpc.DeleteOrderAliAndDb:" >>/tmp/judge14-res.log
      # 转化为字符串
      payPathOrderNumGetOrderDataStr=$(printf '"%d"' "$payPathOrderNumGetOrderDataNum")
      grpcurl -plaintext -d '{"Uuid":'"$uuid"',"UserId":'"$userId"',"PayPathOrderNum":'"$payPathOrderNumGetOrderDataStr"'}' localhost:8010 pb.orderrpc.DeleteOrderAliAndDb >>/tmp/judge14-res.log
      # {"Uuid":1656963502482395136,"UserId":1,"PayPathOrderNum":"1656963502482395136"}
      sleep 1
      #pb.orderrpc.GetOrderInfoByUUIDAndUserId
      #pb.orderrpc.GetOrderInfoByUserIdAndProductId
      #pb.orderrpc.GetOrderStatusByUUID
      #pb.orderrpc.GetOrderUUIdByLimitAndStatus
      #pb.orderrpc.GetProductBindAndPrices
      #pb.orderrpc.GetUserEarnById
      #pb.orderrpc.SearchOrderByCondition
      #pb.orderrpc.SearchUserEarnByCondition
      #pb.orderrpc.UpdateOrder
      #pb.orderrpc.UpdateOrderStatus
      #pb.orderrpc.UpsertUserEarn

      user_res="/tmp/judge14-res.log"
      ans_res="/data/workspace/myshixun/step/step14/judge_ans/ans.log"
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $user_res
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $ans_res
      differAns="$(diff $user_res $ans_res)"
      if [ -z "$differAns" ]; then
        echo "Yes"
        echo "评测结束"
        cleaEnv
        exit
      else
        echo "评测失败"
        echo "测试输出为:"
        cat /tmp/judge14-res.log
        echo "order程序输出为:"
        cat /tmp/judge14-run-order.log
        echo "trade程序输出为:"
        cat /tmp/judge14-run-trade.log
        cleaEnv
        exit
      fi
    fi
  done; then
    cleaEnv
  else
    echo "测评超时"
    cleaEnv
  fi
}

# 启动后台任务执行函数
testing &
# 记录后台任务的pid
job_pid=$!
# 等待后台任务执行完毕
wait $job_pid

# 预期测评结果
#初始化数据库
#mysql: [Warning] Using a password on the command line interface can be insecure.
#开始测评
#编译代码
#编译trade-rpc
#编译order-rpc
#Yes
#评测结束
