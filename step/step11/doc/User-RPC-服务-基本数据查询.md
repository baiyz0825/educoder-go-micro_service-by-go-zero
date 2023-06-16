[TOC]

---
### 任务描述
前面的任务中，我们已经学习了，使用Go-zero框架搭建基本的RPC服务框架结构，以及使用Gorm实现基本的数据库ORM框架接入，下面这个关卡，我们就来实现用户模块的RPC服务

### 相关知识
- Go-Zero框架
- Protobuf文件的编写
- Gorm框架

### 编程要求
#### 提示
注意：本关卡已经帮你生成了对应的RPC接口，你只需要补充RPC接口中的方法(loginc中)逻辑实现即可
- 原始的RPC接口定义为: `sources/apps/user/cmd/rpc/pb/userrpc.proto`
- 原始的GORM生成配置为: `sources/apps/user/gen/generate_test.go`

下面为本关卡具体的项目结构:
```text
├── etc // 配置文件目录
│   └── userrpc.yaml // 配置文件
├── internal // 内部包
│   ├── config // 配置类包
│   │   └── config.go // 配置类
│   ├── dao // Gorm生成的数据库持久层对象包
│   │   ├── major.gen.go // 主修
│   │   ├── major.gen_test.go
│   │   ├── query.go // 查询封装体
│   │   ├── query_test.go
│   │   ├── third.gen.go // 三方
│   │   ├── third.gen_test.go
│   │   ├── third_data.gen.go // 三方数据
│   │   ├── third_data.gen_test.go
│   │   ├── user.gen.go // 用户
│   │   └── user.gen_test.go
│   ├── logic // 实际用户逻辑
│   │   ├── addmajorlogic.go // 增加主修
│   │   ├── addthirddatalogic.go // 增加三方数据
│   │   ├── addthirdlogic.go // 增加三方
│   │   ├── adduserlogic.go // 增加用户
│   │   ├── delmajorlogic.go // 删除主修
│   │   ├── delthirddatalogic.go // 删除三方数据
│   │   ├── delthirdlogic.go // 删除三方
│   │   ├── deluserlogic.go // 删除用户
│   │   ├── getmajorbyidlogic.go // 通过id查询主修
│   │   ├── getmajorpageslogic.go // 分页获取主修
│   │   ├── getthirdbinddatalogic.go // 获取三方数据绑定
│   │   ├── getthirdbyidlogic.go // 通过id获取三方数据
│   │   ├── getthirdbyuseridandtypelogic.go // 通过用户Id获取三方数据
│   │   ├── getthirddatabyidlogic.go // 通过id获取三方数据
│   │   ├── getthirddatabythirdidlogic.go // 通过三方Id获取三方数据
│   │   ├── getuserbyidlogic.go // 通过用户id查询用户
│   │   ├── getuserbyphoneoremaillogic.go // 通过用户手机号或者邮箱查询用户
│   │   ├── updatemajorlogic.go // 更新主修
│   │   ├── updatethirddatalogic.go // 更新三方数据
│   │   ├── updatethirdlogic.go // 更新三方
│   │   └── updateuserlogic.go // 更新用户
│   ├── model // GOrm生成的ORM 结构体
│   │   ├── major.gen.go // 主修
│   │   ├── third.gen.go // 三方
│   │   ├── third_data.gen.go //三方数据
│   │   └── user.gen.go // 用户
│   ├── server // RPC服务器
│   │   └── userrpcserver.go // RPC服务
│   └── svc // 服务上下文
│       └── servicecontext.go // 服务上下文Context
├── pb // Protobuf包
│   ├── userrpc.pb.go // 生成的Go的protobuf结构
│   ├── userrpc.proto // Protobuf原始结构
│   └── userrpc_grpc.pb.go // 生成的Protobuf Grpc基础包
├── userrpc // 请求接口定义
│   └── userrpc.go 
└── userrpc.go // main主程序
```
#### 任务
**请按照下面步骤，修改完成相对应的CURD方法(logic)**
1. 请修改`sources/apps/user/cmd/rpc/etc/userrpc.yaml`配置文件中数据库以及Redis链接如下:

| 类别    | 用户名  | 密码     | 端口   | 地址        |
|-------|------|--------|------|-----------|
| Redis | 空    | 空      | 6879 | 127.0.0.1 |
| Mysql | root | 123123 | 3306 | 127.0.0.1 |

2. 开发实现下列接口逻辑
```text
├── adduserlogic.go 增加用户
├── delmajorlogic.go 删除一个主修课程
├── delthirddatalogic.go 删除第三方数据
├── delthirdlogic.go 删除三方
├── deluserlogic.go 删除用户
├── getmajorbyidlogic.go 通过id获取主修课程
├── getmajorpageslogic.go 通过分页查询所有的主修课程
├── getthirdbinddatalogic.go 获取第三方绑定的数据
├── getthirdbyidlogic.go 通过三方id获取三方数据
├── getthirdbyuseridandtypelogic.go 通过用户id和三方数据类型获取三方数据
├── getthirddatabyidlogic.go 通过Id查询用户三方数据
├── getthirddatabythirdidlogic.go 通过三方获取三方数据
├── getuserbyidlogic.go 通过id获取用户
├── getuserbyphoneoremaillogic.go 通过手机号或者用户名查询用户信息
├── updatemajorlogic.go 更新主修
├── updatethirddatalogic.go 更新三方数据
├── updatethirdlogic.go 更新三方
└── updateuserlogic.go 更新用户信息
```
开发提示:
**接口中入参和出参数，均可在`sources/apps/user/cmd/rpc/pb/userrpc.proto`proto文件中定义找到，编辑器中也可以查看相对应的proto数据类型定义。**<br>
**数据库GO模型，请在`sources/apps/user/cmd/rpc/internal/model`中进行查阅**<br>
**数据上下文设置，请查看`sources/apps/user/cmd/rpc/internal/svc/servicecontext.go`里面进行了数据库、Redis的初始化工作**<br>
**只需要修改logic以及配置文件，其余部分不需要关注**

下面以`adduserlogic.go`方法进行示例:
1. 打开配置文件修改相对应的配置(参照要求)
2. 打开文件`adduserlogic.go`开始编写业务代码
```go
package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ----------------------------------你需要补充的函数内容-----------------------------------
func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// 获取in中获取到的用户信息 具体属性请参照pb.AddUserReq定义
	// Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         //用户名称
	// Password string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` //用户密码
	// Age      int64   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`          //用户年龄
	// Gender   int64   `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`    //用户性别
	// Phone    string  `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`       //用户电话
	// Email    string  `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`       //用户邮箱
	// Grade    int64   `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`      //用户年纪 （大一、大二、大三、大四） 1，2，3，4
	// Major    int64   `protobuf:"varint,8,opt,name=major,proto3" json:"major,omitempty"`      //用户专业信息(关联字段)
	// Star     float64 `protobuf:"fixed64,9,opt,name=star,proto3" json:"star,omitempty"`       //用户等级(0~5)
	// Avatar   string  `protobuf:"bytes,10,opt,name=avatar,proto3" json:"avatar,omitempty"`    //用户头像链接
	// Sign     string  `protobuf:"bytes,11,opt,name=sign,proto3" json:"sign,omitempty"`        //用户个性签名
	// Class    int64   `protobuf:"varint,12,opt,name=class,proto3" json:"class,omitempty"`     //用户班级
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 检查相对应的参数
	if len(in.GetName()) == 0 || len(in.GetPassword()) == 0 || len(in.GetPhone()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 创建上下文对象
	user := &model.User{}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// 从请求数据中获取用户数据
	err := copier.Copy(user, in)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// 初始化用户uuid
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成系统id错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	user.UniqueID = id
	if err := l.svcCtx.Query.User.WithContext(ctx).Create(user); err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	// 成功返回响应结果
	return &pb.AddUserResp{}, nil
}
```
注意：
1. 返回的**错误信息**，请使用`xerr.NewErrCode(xerr.PB_CHECK_ERR)`类似形式返回，具体方法在`sources/common/xerr`包中，**请直接调用**即可。
2. 返回消息中的**常量**`xerr.PB_CHECK_ERR`等，均在`sources/common/xerr/xerr_code.go`中
3. **请在编写业务过程中及时处理返回的错误信息，否则编译时将不会通过**
4. 示例代码中的`copier.Copy(user, in)`方法的使用，请参照[copier文档](https://github.com/jinzhu/copier)中使用，如不使用，也可以使用**手动赋值**的形式实现，如下:

```go
// 声明对象，手动赋值
	user:=&model.User{
		UID:        0,
		UniqueID:   0,
		Name:       "",
		Password:   "",
		Age:        nil,
		Gender:     nil,
		Phone:      "",
		Email:      nil,
		Grade:      nil,
		Major:      nil,
		Star:       nil,
		Avatar:     nil,
		Sign:       nil,
		Class:      nil,
		CreateTime: nil,
		UpdateTime: nil,
		DeleteTime: gorm.DeletedAt{},
	}
```
5. 用户模块，用户`UniqueID`字段请使用**工具包**`utils`**中的方法生成**，如下:
```go
id, err := utils.GenSnowFlakeId()
```
### 测试说明
请在完成**所有接口**之后，在进行评测，平台会评测上述要求中的**部分接口**，**其结果可以满足平台进行增删改查配置的标准输入输出时，视为通过测评，否则不给予通过**
会被评测的接口内容如下:
- pb.userrpc.AddMajor
- pb.userrpc.AddUser
- pb.userrpc.GetMajorPages
- pb.userrpc.GetThirdDataByThirdId
- pb.userrpc.GetUserById
- pb.userrpc.GetUserByPhoneOrEmail
- pb.userrpc.UpdateUser
注：如需自己测试自己服务是否编写正常，可以使用`go run`命令自行运行主程序`sources/apps/user/cmd/rpc/userrpc.go`，或者使用在线编译器进行调试。通过`grpcurl`测试自己的RPC服务接口
> `grpcurl`参考: [掘金-使用 grpcurl 通过命令行访问 gRPC 服务](https://juejin.cn/post/7013612865823178782)
> 查询你服务的RPC接口:`grpcurl -plaintext host:port list` 注意务必开启反射:`reflection.Register(grpcServer)`