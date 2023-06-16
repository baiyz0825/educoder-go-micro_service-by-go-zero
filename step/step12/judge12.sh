#!/bin/bash
source /etc/profile
# 清理环境
cleaEnv() {
  kill $(lsof -t -i:8011) >/dev/null 2>&1
  rm /tmp/judge12-res.log >/dev/null 2>&1
  rm /tmp/judge12-run.log >/dev/null 2>&1
}
cleaEnv
# 定义连接参数
MYSQL_USER="root"
MYSQL_PASS="123123"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="3306"
RPC_URL="localhost:8011"
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
if mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "SHOW DATABASES LIKE 'school_resources';" | grep -q school_resources; then
  mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" -e "DROP DATABASE school_resources;" >/dev/null 2>&1
fi
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step12/judge_data/school_resources.sql >/dev/null 2>&1
echo "开始测评"
touch /tmp/judge12-run.log
touch /tmp/judge12-res.log
cd /data/workspace/myshixun/sources/apps/resources/cmd/rpc || exit
echo "编译代码"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/resources/cmd/rpc/resourcesrpc.go' >/tmp/judge12-run.log 2>&1 &
# 等待服务启动
sleep 8

testing() {
  # timeout命令用来在一定时间内执行tail -f命令，如果超时就会结束循环
  if timeout 240 tail -f /tmp/judge12-run.log | while read line; do
    # 检查关键词
    if echo "$line" | grep -q "Start"; then
      # pb.resourcesrpc.SearchClassificationAll
      echo -e "评测pb.resourcesrpc.SearchClassificationAll:" >>/tmp/judge12-res.log
      grpcurl -plaintext localhost:8011 pb.resourcesrpc.SearchClassificationAll >>/tmp/judge12-res.log
      sleep 1
      # pb.resourcesrpc.GetClassificationDataByPages
      echo -e "评测pb.resourcesrpc.GetClassificationDataByPages:" >>/tmp/judge12-res.log
      grpcurl -plaintext -d '{"page": 1,"limit": 200,"classificationID": 1}' localhost:8011 pb.resourcesrpc.GetClassificationDataByPages >>/tmp/judge12-res.log
      sleep 1
      # pb.resourcesrpc.AddOnlineText
      echo -e "评测pb.resourcesrpc.AddOnlineText:" >>/tmp/judge12-res.log
      grpcurl -plaintext -d '{"typeSuffix":0,"owner":1,"content":"这是在线文本测试数据1","classID":4,"permission":1,"textPoster":"https://www.bai/post.img","textName":"文本测试1"}' localhost:8011 pb.resourcesrpc.AddOnlineText >>/tmp/judge12-res.log
      sleep 1
      # pb.resourcesrpc.UpdateFile
      echo -e "评测pb.resourcesrpc.UpdateFile:" >>/tmp/judge12-res.log
      grpcurl -plaintext -d '{"ID":14,"name":"艺术体操Mod","obfuscateName":"xascaca","size":2271,"owner":2,"status":1,"type":0,"class":2,"suffix":".zip","downloadAllow":1,"link":"resources/files/艺术体操.zip","filePoster":"https://this_is.com/poster.img"}' localhost:8011 pb.resourcesrpc.UpdateFile >>/tmp/judge12-res.log
      sleep 1
      # pb.resourcesrpc.DelCount
      echo -e "评测pb.resourcesrpc.DelCount:" >>/tmp/judge12-res.log
      grpcurl -plaintext -d '{"id":6}' localhost:8011 pb.resourcesrpc.DelCount >>/tmp/judge12-res.log
      sleep 1
      # pb.resourcesrpc.DelFile
      echo -e "评测pb.resourcesrpc.DelFile:" >>/tmp/judge12-res.log
      grpcurl -plaintext -d '{"ID":14}' localhost:8011 pb.resourcesrpc.DelFile >>/tmp/judge12-res.log
      # 输出结果 比对答案
      user_res="/tmp/judge12-res.log"
      ans_res="/data/workspace/myshixun/step/step12/judge_ans/ans.log"
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
        cat /tmp/judge12-res.log
        echo "程序输出为:"
        cat /tmp/judge12-run.log
        cleaEnv
        exit
      fi
    fi
  done; then
    echo "------------"
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
# 测评成功输出
#mysql: [Warning] Using a password on the command line interface can be insecure.
#开始测评
#编译代码
#Yes
#评测结束
#------------
