#!/bin/bash
source /etc/profile
# 清理环境
cleaEnv() {
  kill $(lsof -t -i:8012) >/dev/null 2>&1
  rm /tmp/judge13-res.log >/dev/null 2>&1
  rm /tmp/judge13-run.log >/dev/null 2>&1
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
if mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "SHOW DATABASES LIKE 'school_trade';" | grep -q school_trade; then
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_trade;" >/dev/null 2>&1
fi
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step13/judge_data/school_trade.sql >/dev/null 2>&1
echo "开始测评"
touch /tmp/judge13-run.log
touch /tmp/judge13-res.log
cd /data/workspace/myshixun/sources/apps/trade/cmd/rpc || exit
echo "编译代码"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/trade/cmd/rpc/traderpc.go' >/tmp/judge13-run.log 2>&1 &
# 等待服务启动
testing() {
  # timeout命令用来在一定时间内执行tail -f命令，如果超时就会结束循环
  if timeout 240 tail -f /tmp/judge13-run.log | while read line; do
    if echo "$line" | grep -q "Start"; then
      # pb.traderpc.AddProduct
      echo -e "评测pb.traderpc.AddProduct:" >>/tmp/judge13-res.log
      grpcurl -plaintext -d '{"name":"人文的力量中马里奥动画","type":4,"productBind":16,"owner":20,"price":1.2175,"productPoster":"https://abcd/abcd.poster"}' localhost:8012 pb.traderpc.AddProduct >>/tmp/judge13-res.log

      # pb.traderpc.SearchProductByResourcesBind
      echo -e "评测pb.traderpc.SearchProductByResourcesBind:" >>/tmp/judge13-res.log
      SearchProductByResourcesBind=$(grpcurl -plaintext -d '{"resourceId":16}' localhost:8012 pb.traderpc.SearchProductByResourcesBind)
      echo "$SearchProductByResourcesBind" >/tmp/judge13-res-temp.log
      product_name=$(echo "$SearchProductByResourcesBind" | grep -Po '"productName"\s*:\s*"\K[^"]*')
      echo "$product_name" >>/tmp/judge13-res.log

      # pb.traderpc.GetProductById
      # 过滤出产品id数字
      productId=$(grep 'productId' /tmp/judge13-res-temp.log | cut -d":" -f2 | tr -d ' "' | sed 's/,//g; s/ //g')
      rm /tmp/judge13-res-temp.log
      echo -e "评测pb.traderpc.GetProductById:" >>/tmp/judge13-res.log
      getProductByIdResp=$(grpcurl -plaintext -d '{"ID":'"$productId"',"UUID":0}' localhost:8012 pb.traderpc.GetProductById)
      name=$(echo "$getProductByIdResp" | grep -Po '"name"\s*:\s*"\K[^"]*')
      echo "$name" >>/tmp/judge13-res.log

      # pb.traderpc.SearchProduct
      echo -e "评测pb.traderpc.SearchProduct:" >>/tmp/judge13-res.log
      searchProductResp=$(grpcurl -plaintext -d '{"page": 1,"limit": 100,"type": 1,"bottomPrice": 1,"topPrice": 2.21,"desc": true}' localhost:8012 pb.traderpc.SearchProduct)
      productPoster=$(echo "$searchProductResp" | grep -Po '"productPoster"\s*:\s*"\K[^"]*')
      echo "$productPoster" >>/tmp/judge13-res.log

      user_res="/tmp/judge13-res.log"
      ans_res="/data/workspace/myshixun/step/step13/judge_ans/ans.log"
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
        cat /tmp/judge13-res.log
        echo "程序输出为:"
        cat /tmp/judge13-run.log
        cleaEnv
        exit
      fi
    fi
  done; then
    echo "--------------"
  else
    echo "评测超时"
    cleaEnv
  fi
}

# 启动后台任务执行函数
testing &
# 记录后台任务的pid
job_pid=$!
# 等待后台任务执行完毕
wait $job_pid
# 评测正确输出结果
#mysql: [Warning] Using a password on the command line interface can be insecure.
#开始测评
#编译代码
#Yes
#评测结束
#--------------
