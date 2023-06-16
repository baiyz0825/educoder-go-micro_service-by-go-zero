#!/bin/bash
#----------------------------------------------------测试主程序是否正常运行----------------------------------------------------
source /etc/profile
kill  $(lsof -t -i:8889)  > /dev/null 2>&1
# 运行 user.go 程序
touch /tmp/judge3-run.log
cd  /data/workspace/myshixun/step/step3/stu || exit
nohup bash -c 'go run /data/workspace/myshixun/step/step3/stu/user.go' > /tmp/judge3-run.log 2>&1 &
# 等待程序启动，可以根据实际情况调整等待时间
sleep 8
# 发送 post 请求并输出响应
response=$(curl -s -X POST -H "Content-Type: application/json" -d '{"userName":"abcd","passwd":"123avcd"}' -g http://localhost:8889/login)
# 判断失败
if [[ "$response" != '{"token":"abcd123avcd"}' ]]; then
  cat /tmp/judge3-run.log
fi
# 结束 user.go 程序 输出信息重定向到错误
kill  $(lsof -t -i:8889)  > /dev/null 2>&1
echo "评测成功"
echo "$response"
# 平台预期结果
#评测成功
#{"token":"abcd123avcd"}
