#!/bin/bash
source /etc/profile
# 清理环境
cleaEnv() {
  kill $(lsof -t -i:8013) >/dev/null 2>&1
  rm /tmp/judge11-res.log >/dev/null 2>&1
  rm /tmp/judge11-run.log >/dev/null 2>&1
}
cleaEnv
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
mysql -u"${MYSQL_USER}" -p"${MYSQL_PASS}" -h"${MYSQL_HOST}" -P"${MYSQL_PORT}" </data/workspace/myshixun/step/step11/judge_data/school_user.sql >/dev/null 2>&1
echo "开始测评"
# 创建日志文件
touch /tmp/judge11-run.log
touch /tmp/judge11-res.log
# 编译启动服务
cd /data/workspace/myshixun/sources/apps/user/cmd/rpc || exit
echo "编译代码"
nohup sh -c 'go run /data/workspace/myshixun/sources/apps/user/cmd/rpc/userrpc.go' >/tmp/judge11-run.log 2>&1 &
# 等待服务启动
testing() {
  # timeout命令用来在一定时间内执行tail -f命令，如果超时就会结束循环
  if timeout 240 tail -f /tmp/judge11-run.log | while read line; do
    # 检查关键词
    if echo "$line" | grep -q "Start"; then
      # 开始进行接口请求
      echo "开始进行接口请求"
      #pb.userrpc.AddMajor
      echo -e "评测pb.userrpc.AddMajor:" >/tmp/judge11-res.log
      grpcurl -plaintext -d '{"name":"人工数据分析","desc":"这是测试科目人工数据分析"}' localhost:8013 pb.userrpc.AddMajor >>/tmp/judge11-res.log
      sleep 1
      # 响应数据
      #{
      #
      #}
      #pb.userrpc.AddThird
      #pb.userrpc.AddThirdData
      #pb.userrpc.AddUser
      echo -e "评测pb.userrpc.AddUser:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"name": "Test3","password": "这是测试数据3","age": 22,"gender": 11,"phone": "19993655016","email": "jrbhwtc4381@xdrjeg.com","grade": 2,"major": 4,"star": 4.7812,"avatar": "https://xxxxx/img","sign": "这是测试数据2","class": 121}' localhost:8013 pb.userrpc.AddUser >>/tmp/judge11-res.log
      sleep 1
      # 响应数据
      #{
      #
      #}
      #pb.userrpc.GetMajorById
      #pb.userrpc.GetMajorPages
      echo -e "评测pb.userrpc.GetMajorPages:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"page":1,"limit":20}' localhost:8013 pb.userrpc.GetMajorPages >>/tmp/judge11-res.log
      sleep 1
      # 响应数据
      #{
      #  "major": [
      #    {
      #      "ID": "1",
      #      "name": "计算机科学",
      #      "desc": "这是升级版本极客"
      #    },
      #    {
      #      "ID": "3",
      #      "name": "Java语言",
      #      "desc": "计算机"
      #    },
      #    {
      #      "ID": "5",
      #      "name": "PHP语言",
      #      "desc": "更新数据"
      #    },
      #    {
      #      "ID": "6",
      #      "name": "Py语言",
      #      "desc": "编 程"
      #    },
      #    {
      #      "ID": "7",
      #      "name": "Rust语言",
      #      "desc": "这是升级版本极客"
      #    },
      #    {
      #      "ID": "8",
      #      "name": "R语言",
      #      "desc": "新语言"
      #    },
      #    {
      #      "ID": "9",
      #      "name": "美术",
      #      "desc": "分类"
      #    },
      #    {
      #      "ID": "10",
      #      "name": "科技",
      #      "desc": "分类"
      #    },
      #    {
      #      "ID": "11",
      #      "name": "人工数据分析",
      #      "desc": "这是测试科目人工数据分析"
      #    }
      #  ]
      #}

      #pb.userrpc.GetThirdBindData
      #pb.userrpc.GetThirdById
      #pb.userrpc.GetThirdByUserIdAndType
      #pb.userrpc.GetThirdDataById
      #grpcurl -plaintext -d '{"ID":2}' localhost:8013  pb.userrpc.GetThirdDataById
      # 测试数据
      #{
      #  "thirdData": {
      #    "ID": "2",
      #    "thirdID": "2",
      #    "name": "Test2Gitee",
      #    "sign": "这是测试数据2的Gitee"
      #  }
      #}
      #pb.userrpc.GetThirdDataByThirdId
      echo -e "评测pb.userrpc.GetThirdDataByThirdId:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"thirdID":1}' localhost:8013 pb.userrpc.GetThirdDataByThirdId >>/tmp/judge11-res.log
      sleep 1
      # 测试数据
      #{
      #  "thirdData": {
      #    "ID": "1",
      #    "thirdID": "1",
      #    "name": "Test1QQ",
      #    "sign": "这是测试数据1的qq"
      #  }
      #}
      #pb.userrpc.GetUserById
      echo -e "评测pb.userrpc.GetUserById:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"ID":1}' localhost:8013 pb.userrpc.GetUserById >>/tmp/judge11-res.log
      sleep 1
      # 测试数据
      #{
      #  "user": {
      #    "uID": "1",
      #    "uniqueID": "1653421638231789568",
      #    "name": "test1",
      #    "password": "test1234",
      #    "age": "22",
      #    "gender": "1",
      #    "phone": "18178114924",
      #    "email": "ovgtmcffm@163.com",
      #    "grade": "1",
      #    "major": "1",
      #    "star": 2,
      #    "avatar": "https://abcd.asas/img",
      #    "sign": "这是测试数据一",
      #    "class": "192",
      #    "createTime": "1683772307000",
      #    "updateTime": "1683772397000"
      #  }
      #}
      # -------------------------------------------------------------------------
      #pb.userrpc.GetUserByPhoneOrEmail
      echo -e "评测pb.userrpc.GetUserByPhoneOrEmail:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"phone":"18178114924","email":"ovgtmcffm@163.com"}' localhost:8013 pb.userrpc.GetUserByPhoneOrEmail >>/tmp/judge11-res.log
      sleep 1
      # 测试数据
      #{
      #  "user": {
      #    "uID": "1",
      #    "uniqueID": "1653421638231789568",
      #    "name": "test1",
      #    "password": "test1234",
      #    "age": "22",
      #    "gender": "1",
      #    "phone": "18178114924",
      #    "email": "ovgtmcffm@163.com",
      #    "grade": "1",
      #    "major": "1",
      #    "star": 2,
      #    "avatar": "https://abcd.asas/img",
      #    "sign": "这是测试数据一",
      #    "class": "192",
      #    "createTime": "1683772307000",
      #    "updateTime": "1683772397000"
      #  }
      #}
      #pb.userrpc.DelMajor
      #pb.userrpc.DelThird
      #pb.userrpc.DelThirdData
      #pb.userrpc.DelUser

      #pb.userrpc.UpdateMajor
      #pb.userrpc.UpdateThird
      #pb.userrpc.UpdateThirdData
      #pb.userrpc.UpdateUser
      echo -e "评测pb.userrpc.UpdateUser:" >>/tmp/judge11-res.log
      grpcurl -plaintext -d '{"uID": 1,"uniqueID": 1653421638231789568,"name": "test1-mod","password": "test1234aa","age": 33,"gender": 22,"phone": "18178114924","email": "ovgtmcfssfm@163.com","grade": 2,"major": 3,"star": 7.871,"avatar": "","sign": "更新数据之后的用户信息","class": 192}' localhost:8013 pb.userrpc.UpdateUser >>/tmp/judge11-res.log
      sleep 1
      # 响应输出
      #{
      #
      #}
      # 输出结果 比对答案
      user_res="/tmp/judge11-res.log"
      ans_res="/data/workspace/myshixun/step/step11/judge_ans/ans.log"
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
        cat /tmp/judge11-res.log
        #        echo "程序输出为:"
        #        cat /tmp/judge11-run.log
        echo "评测结束"
        cleaEnv
        exit
      fi
    fi
  done; then
    echo "-----------------"
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

#正确测评结果:
#mysql: [Warning] Using a password on the command line interface can be insecure.
#开始测评
#编译代码
#开始进行接口请求
#Yes
#评测结束
#-----------------
