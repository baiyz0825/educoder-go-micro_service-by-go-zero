[TOC]

---
### 任务描述
前面的任务中，我们已经学习了，使用Go-zero框架搭建基本的RPC服务框架结构，以及使用Gorm实现基本的数据库ORM框架接入，下面这个关卡，我们就来实现商品模块的RPC服务

### 相关知识
- Go-Zero框架
- Protobuf文件的编写
- Gorm框架

### 编程要求
#### 提示
注意：本关卡已经帮你生成了对应的RPC接口，你只需要补充RPC接口中的方法(loginc中)逻辑实现即可
- 原始的RPC接口定义为: `sources/apps/trade/cmd/rpc/pb/traderpc.proto`
- 原始的GORM生成配置为: `sources/apps/trade/gen/generate_test.go`

下面为本关卡具体的项目结构:
```text
├── etc
│   └── traderpc.yaml
├── internal
│   ├── config
│   │   └── config.go
│   ├── dao
│   │   ├── product.gen.go
│   │   ├── product.gen_test.go
│   │   ├── query.go
│   │   └── query_test.go
│   ├── logic
│   │   ├── addproductlogic.go
│   │   ├── delproductlogic.go
│   │   ├── getproductbindbyproductidlogic.go
│   │   ├── getproductbybindidandownerlogic.go
│   │   ├── getproductbyidlogic.go
│   │   ├── getproductidandproductnamelogic.go
│   │   ├── searchproductbyresourcesbindlogic.go
│   │   ├── searchproductlogic.go
│   │   └── updateproductlogic.go
│   ├── model
│   │   └── product.gen.go
│   ├── server
│   │   └── traderpcserver.go
│   └── svc
│       └── servicecontext.go
├── pb
│   ├── traderpc.pb.go
│   ├── traderpc.proto
│   └── traderpc_grpc.pb.go
├── traderpc
│   └── traderpc.go
└── traderpc.go
```
#### 任务
**请按照下面步骤，修改完成相对应的CURD方法(logic)**
1. 请修改`sources/apps/trade/cmd/rpc/etc/traderpc.yaml`配置文件中数据库以及Redis链接如下:

| 类别    | 用户名  | 密码     | 端口   | 地址        |
|-------|------|--------|------|-----------|
| Redis | 空    | 空      | 6879 | 127.0.0.1 |
| Mysql | root | 123123 | 3306 | 127.0.0.1 |

2. 开发实现下列接口逻辑
```text
├── addproductlogic.go
├── delproductlogic.go
├── getproductbindbyproductidlogic.go
├── getproductbybindidandownerlogic.go
├── getproductbyidlogic.go
├── getproductidandproductnamelogic.go
├── searchproductbyresourcesbindlogic.go
├── searchproductlogic.go
└── updateproductlogic.go
```
开发提示:
**接口中入参和出参数，均可在`sources/apps/trade/cmd/rpc/pb/traderpc.proto`proto文件中定义找到，编辑器中也可以查看相对应的proto数据类型定义。**<br>
**数据库GO模型，请在`sources/apps/trade/cmd/rpc/internal/model`中进行查阅**<br>
**数据上下文设置，请查看`sources/apps/trade/cmd/rpc/internal/svc/servicecontext.go`里面进行了数据库、Redis的初始化工作**<br>
**只需要修改logic以及配置文件，其余部分不需要关注**
### 测试说明
请在完成**所有接口**之后，在进行评测，平台会评测上述要求中的**部分接口**，**其结果可以满足平台进行增删改查配置的标准输入输出时，视为通过测评，否则不给予通过**
会被评测的接口内容如下:
- pb.traderpc.AddProduct 增加商品
- pb.traderpc.SearchProductByResourcesBind 使用商品绑定的资源id获取相应的商品id
- pb.traderpc.GetProductById 通过商品id获取商品
- pb.traderpc.SearchProduct 检索商品
注：如需自己测试自己服务是否编写正常，可以使用`go run`命令自行运行主程序`sources/apps/trade/cmd/rpc/traderpc.go`，或者使用在线编译器进行调试。通过`grpcurl`测试自己的RPC服务接口
> `grpcurl`参考: [掘金-使用 grpcurl 通过命令行访问 gRPC 服务](https://juejin.cn/post/7013612865823178782)
> 查询你服务的RPC接口:`grpcurl -plaintext host:port list` 注意务必开启反射:`reflection.Register(grpcServer)`
