[TOC]

---
### 任务描述
前面的任务中，我们已经学习了，使用Go-zero框架搭建基本的RPC服务框架结构，以及使用Gorm实现基本的数据库ORM框架接入，下面这个关卡，我们就来实现资源模块的RPC服务

### 相关知识
- Go-Zero框架
- Protobuf文件的编写
- Gorm框架


### 编程要求
#### 提示
注意：本关卡已经帮你生成了对应的RPC接口，你只需要补充RPC接口中的方法(loginc中)逻辑实现即可
- 原始RPC接口定义为: `sources/apps/resources/cmd/rpc/pb/resourcesrpc.proto`
- 原始GORM生成的配置为: `sources/apps/resources/gen/generate_test.go`

下面为本关卡具体的项目结构:
```text
├── etc
│   └── resourcesrpc.yaml
├── internal
│   ├── config
│   │   └── config.go
│   ├── dao
│   │   ├── classification.gen.go
│   │   ├── classification.gen_test.go
│   │   ├── count.gen.go
│   │   ├── count.gen_test.go
│   │   ├── file.gen.go
│   │   ├── file.gen_test.go
│   │   ├── online_text.gen.go
│   │   ├── online_text.gen_test.go
│   │   ├── query.go
│   │   ├── query_test.go
│   │   ├── res_comment.gen.go
│   │   └── res_comment.gen_test.go
│   ├── logic
│   │   ├── addclassificationlogic.go
│   │   ├── addcountlogic.go
│   │   ├── addfilelogic.go
│   │   ├── addonlinetextlogic.go
│   │   ├── addrescommentlogic.go
│   │   ├── checkdownloadallowlogic.go
│   │   ├── delclassificationlogic.go
│   │   ├── delcountlogic.go
│   │   ├── delfilelogic.go
│   │   ├── delonlinetextlogic.go
│   │   ├── delrescommentlogic.go
│   │   ├── getclassificationbyidlogic.go
│   │   ├── getclassificationdatabypageslogic.go
│   │   ├── getcountbyidlogic.go
│   │   ├── getcountbyuidlogic.go
│   │   ├── getfilebyidlogic.go
│   │   ├── getfilresourcesuseridlogic.go
│   │   ├── getonlinetextbyidlogic.go
│   │   ├── getrescommentbyidlogic.go
│   │   ├── searchclassificationalllogic.go
│   │   ├── searchfileconditionpageslogic.go
│   │   ├── searchonlineconditiontextlogic.go
│   │   ├── searchrescommentbyuserorresidlogic.go
│   │   ├── updateclassificationlogic.go
│   │   ├── updatecountlogic.go
│   │   ├── updatefilelogic.go
│   │   ├── updateonlinetextlogic.go
│   │   └── updaterescommentlogic.go
│   ├── model
│   │   ├── classification.gen.go
│   │   ├── count.gen.go
│   │   ├── file.gen.go
│   │   ├── online_text.gen.go
│   │   └── res_comment.gen.go
│   ├── server
│   │   └── resourcesrpcserver.go
│   └── svc
│       └── servicecontext.go
├── pb
│   ├── resourcesrpc.pb.go
│   ├── resourcesrpc.proto
│   └── resourcesrpc_grpc.pb.go
├── resourcesrpc
│   └── resourcesrpc.go
└── resourcesrpc.go
```
#### 任务
1. 请修改`sources/apps/resources/cmd/rpc/etc/resourcesrpc.yaml`配置文件中数据库以及Redis链接如下:

| 类别    | 用户名  | 密码     | 端口   | 地址        |
|-------|------|--------|------|-----------|
| Redis | 空    | 空      | 6879 | 127.0.0.1 |
| Mysql | root | 123123 | 3306 | 127.0.0.1 |
2. 开发实现下列接口逻辑
```text
├── addclassificationlogic.go // 增加分类
├── addcountlogic.go // 增加用户数据统计
├── addfilelogic.go // 增加文件
├── addonlinetextlogic.go // 增加在线文本
├── addrescommentlogic.go // 增加文件资源评价
├── checkdownloadallowlogic.go // 检查资源是否允许下载
├── delclassificationlogic.go // 删除分类
├── delcountlogic.go // 删除用户数据统计
├── delfilelogic.go // 删除文件
├── delonlinetextlogic.go // 删除在线文本数据
├── delrescommentlogic.go // 删除资源评论
├── getclassificationbyidlogic.go // 通过id获取一个分类数据
├── getclassificationdatabypageslogic.go // 分页获取分类下资源的数据(多条件查询)
├── getcountbyidlogic.go // 通过id获取用户资料统计
├── getcountbyuidlogic.go // 通过用户id获取用户的资料数据统计
├── getfilebyidlogic.go // 通过id获取文件
├── getfilresourcesuseridlogic.go // 通过资源id查询资源owner
├── getonlinetextbyidlogic.go // 通过在线文本id查询在线文本数据
├── getrescommentbyidlogic.go // 通过id查询资源评价数据
├── searchclassificationalllogic.go // 获取全部分分类信息（树形列表，没有根节点）
├── searchfileconditionpageslogic.go // 通过条件分页搜索文件
├── searchonlineconditiontextlogic.go // 搜索在线文本资料
├── searchrescommentbyuserorresidlogic.go //查询某一个资源下评论  或者 查询用户全部评论
├── updateclassificationlogic.go // 更新分类
├── updatecountlogic.go // 更新用户评分
├── updatefilelogic.go // 更新文件资源
├── updateonlinetextlogic.go // 更新在线文本
└── updaterescommentlogic.go // 更新文本资源评价
```
开发提示:
**接口中入参和出参数，均可在`sources/apps/resources/cmd/rpc/pb/resourcesrpc.proto`proto文件中定义找到，编辑器中也可以查看相对应的proto数据类型定义。**<br>
**数据库GO模型，请在`sources/apps/resources/cmd/rpc/internal/model`中进行查阅**<br>
**数据上下文设置，请查看`sources/apps/resources/cmd/rpc/internal/svc/servicecontext.go`里面进行了数据库、Redis的初始化工作**<br>
**只需要修改logic以及配置文件，其余部分不需要关注**<br>
**具体编写方法请参照前一张任务示例**
### 测试说明
请在完成**所有接口**之后，在进行评测，平台会评测上述要求中的**部分接口**，**其结果可以满足平台进行增删改查配置的标准输入输出时，视为通过测评，否则不给予通过**
会被评测的接口内容如下:
- pb.resourcesrpc.SearchClassificationAll: 这里获取的是一个无根节点的树形菜单数据
- pb.resourcesrpc.GetClassificationDataByPages: 这里存在多条件查询，分页参数必传，根据文件类型(xonst.FIlE_TYPE、xconst.TEXT_TYPE)。其次需要检查是否传递了，分类id、用户id关键词，如果传递请拼接查询条件
- pb.resourcesrpc.AddOnlineText
- pb.resourcesrpc.UpdateFile
- pb.resourcesrpc.DelCount
- pb.resourcesrpc.DelFile

注：如需自己测试自己服务是否编写正常，可以使用`go run`命令自行运行主程序`sources/apps/resources/cmd/rpc/resourcesrpc.go`，或者使用在线编译器进行调试。通过`grpcurl`测试自己的RPC服务接口
> `grpcurl`参考: [掘金-使用 grpcurl 通过命令行访问 gRPC 服务](https://juejin.cn/post/7013612865823178782)
> 查询你服务的RPC接口:`grpcurl -plaintext host:port list` 注意务必开启反射:`reflection.Register(grpcServer)`
