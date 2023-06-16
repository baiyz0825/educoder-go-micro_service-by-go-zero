[TOC]
---
### 任务描述
本关卡任务：编写一个符合Go-zero的RPC文件，并生成想对应的RPC接口

### 相关知识
1. Goctl:[Goctl-rpc](https://go-zero.dev/cn/docs/advance/rpc-call/)

#### RPC文件编写规则
Go-Zero的RPC基于GRPC编写，GRPC服务的创建依赖于Protobuf文件，在上一节我们已经学习了如何编写Protobuf文件。

#### RPC文件生成结构
由于框架提供了一套定义API文件的基本语法，因此直接使用框架工具进行生成相对应的基本桩代码: `goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.`<br>
生成之后的目录文件如下:
```txt
.
├── etc
│   └── user.yaml // 配置文件
├── go.mod // go模块依赖管理
├── internal // 内部包
│   ├── config
│   │   └── config.go //配置类
│   ├── logic // 实际接口业务逻辑
│   │   └── getuserinfologic.go
│   ├── server // 定义的rpc服务
│   │   └── orderserviceserver.go
│   └── svc // 服务上下文
│       └── servicecontext.go
├── orderservice // 实际的service
│   └── orderservice.go
├── types // 生成的pb golang文件
│   └── pb
│       ├── user.pb.go
│       └── user_grpc.pb.go
├── user.go // 主程序
└── user.proto // protobuf描述文件
```

### 编程要求
> 请不要更换修改文件的名称，否则会影响程序判定正确性，你需要修改的文件是：step/step5/stu/user.proto
请按照下面要求编写,protobuf，文件并生成相应的RPC框架桩代码
1. Proto文件中包含两个Message
- UserReq:
    - UserName(string)
    - UserId(int64)两个参数
- UserResp:
    - UserHome(string)
    - UserMoneyCardNum(int64): 注意这个UserMoneyCardNum存在多个。
2. 定义一个RPC方法，方法签名是，`GetUserInfo`入参是UserReq，返回值是UserResp。
3. 使用goctl生成对应的RPC服务（请进入目录`step/step5/stu`）执行生成桩代码
4. 编写成功生成的桩代码中的(`step/step5/stu/internal/logic/getuserinfologic.go`)文件编写响应的逻辑代码，通过将我发送的请求参数中
Passwd 和 UserName 进行拼接，放入UserHome中，之后UserHome对应的UserMoneyCardNum为Passwd重复三次
5. 修改配置文件(`step/step5/stu/etc/user.yaml`)文件，中系统运行端口为`8888`,删除`ETCD`部分配置，增加`Mode: dev`
完整示例如下:
```yaml
Name: user.rpc
ListenOn: 0.0.0.0:8888
Mode: dev
```
### 测试说明
平台会按照上述示例发送请求，测试是否符合上述预期逻辑输出，如果符合则通过，不符合则不通过
#### 示例
请求:
```json
{
  "UserName": "abcd",
  "Passwd": "1234"
}
```
响应数据:
```json
{
  "UserHome":"abcd1234",
  "UserMoneyCardNum":[
    "1234","1234","1234"
  ]
}
```
---
开始你的任务吧，祝你成功！