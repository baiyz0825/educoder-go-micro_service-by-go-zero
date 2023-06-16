[TOC]

---
### 任务描述
在本关卡，我们将编写Resources-API服务，这个服务主要负责将RPC中的数据转化为API接口形式，并且加上相关权限检查，返回。其次在本服务中我们将，与阿里云OSS进行对接，将文件传输存储在云端。
### 相关知识
- JSON
- Golang
- Protobuf
- Go-Zero
- Golang Routine


### 编程要求
#### 提示
1. 请修改对应的配置文件，配置对应的Rpc连接地址(如是按照前面关卡要求完成的编写，只需要检查此部分`sources/apps/trade/cmd/api/etc/trade.yaml`配置是否与下述一致即可):
```yaml
TradeRpc:
  Endpoints:
    - 0.0.0.0:8012
ResourcesRpc:
  EndPoints:
    - 0.0.0.0:8011
```
5. 配置自己的阿里云OSS对应的密钥
```yaml
AliCloud:
  AccessKeyId: 
  AccessKeySecret: 
  BucketName: 
  EndPoint:
  CommonPath:
```
#### 任务
1. 提示
编写`addproductlogic`逻辑的时候，需要检查商品是否以及存在，一个用户发布的file资源只能被自己上架一次！！
2. 完成所有login下的API接口逻辑
```txt
└── product
    ├── addproductlogic.go // 增加商品信息
    ├── deloneproductlogic.go // 删除商品
    ├── getproductinfologic.go // 查询商品详情
    └── getproductinfoquerylogic.go // 条件查询商品详情
```

### 测试说明
测试前提: 平台会编译前面任务完成的`user-rpc`、`resource-rpc`,`user-api`,`trade-rpc`服务，之后启动当前关卡的服务进行测评<br>
请在完成**所有接口**之后，在进行评测，平台会评测上述要求中的**部分接口**，**其结果可以满足平台进行增删改查配置的标准输入输出时，视为通过测评，否则不给予通过**
会被评测的接口内容如下:
- `/product/addone`
- `/product/search`