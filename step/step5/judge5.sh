#!/bin/bash
#----------------------------------------------------测试主程序是否正常运行----------------------------------------------------
source /etc/profile
# 清理环境
cleaEnv() {
  # 结束 user.go 程序 输出信息重定向到错误
  kill $(lsof -t -i:8888) >/dev/null 2>&1
  # 删除缓存
  rm -f /tmp/judge5-res.log >/dev/null 2>&1
  rm /tmp/judge5-run.log >/dev/null 2>&1
}
cleaEnv
# 运行 user.go 程序
touch /tmp/judge5-run.log
touch /tmp/judge5-res.log
cd /data/workspace/myshixun/step/step5/stu || exit
nohup bash -c 'go run /data/workspace/myshixun/step/step5/stu/user.go' >/tmp/judge5-run.log 2>&1 &
# 监听 /tmp/judge5-run.log 文件的变化
testing() {
  # timeout命令用来在一定时间内执行tail -f命令，如果超时就会结束循环
  if timeout 240 tail -f /tmp/judge5-run.log | while read line; do
    # 检查关键词
    if echo "$line" | grep -q "Start"; then
      # 发送 rpc 请求并输出响应
      response=$(grpcurl -plaintext -d '{"UserName": "abcd", "Passwd": 123123}' localhost:8888 pb.OrderService.GetUserInfo)
      echo "$response" >/tmp/judge5-res.log
      # 去除空格进行比较
      user_res="/tmp/judge5-res.log"
      ans_res="/data/workspace/myshixun/step/step5/judge_ans/ans.log"
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $user_res
      sed -i '/^\s*$/d;s/[[:space:]]\+//g' $ans_res
      differAns="$(diff $user_res $ans_res)"
      if [ -z "$differAns" ]; then
        # 实际请求响应符合预期，退出整个脚本
        echo "Yes!"
        cleaEnv
        # 退出整个脚本
        exit
      else
        # 实际请求响应不符合预期，退出整个脚本
        echo -e "程序运行状态: "
        cat /tmp/judge5-run.log
        echo -e "实际请求响应: "
        cat /tmp/judge5-res.log
        cleaEnv
        # 退出整个脚本
        exit
      fi
    fi
  done; then
    cleaEnv
  else
    echo "测评超时!"
    cleaEnv
  fi
}
# 启动后台任务执行函数
testing &
# 记录后台任务的pid
job_pid=$!
# 等待后台任务执行完毕
wait $job_pid

# 平台预期结果
# Yes!
