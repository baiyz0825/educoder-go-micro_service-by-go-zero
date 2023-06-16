# 介绍

[TOC]

---

### 任务描述
本关卡任务：编写一个符合Go-zero的API文件，并生成对应的API接口

### 相关知识
1. 如何编写API文件 : https://go-zero.dev/cn/docs/advance/api-coding
2. 如何生成基本API文件: https://go-zero.dev/cn/docs/goctl/api

#### API文件编写规则

一个符合Go-Zero的API文件基本分为以下几个部分: 
1. 语法声明: 这部分主要描述API文件的基本语法版本<br>
example: `syntax = "v1"`
2. 导入模块: 方便接口之间数据定义和接口方法解耦，这里可以使用导入，导入其他API文件中的结构

```api
// 从其他文件导入单个API文件，如不需导入可以不写
import "foo.api"

// 从其他文件导入多个API文件，如不需导入可以不写
import (
    "bar.api"
    "foo/bar.api"
)
```
> 这里，需要注意的是，包含下述Service模块的API文件只能有一个（主文件），因此一般导入的为其他API文件中定义的**结构**
3. API文件描述: 这里的描述，会在通过工具生成对应的Swagger(接口文档)的时候成为这一组API的整体描述信息，具体内容如下:

```api
// API文件描述
info(
    author: "songmeizi" // 作者
    date:   "2020-01-08" // 日期
    desc:   "api语法示例及语法说明" // 文档描述
)
```
4. 数据声明: 数据声明部分，通过近似Golang语言基本语法，声明对应的数据结构，方便框架对请求参数、请求体、响应体进行对应的封装。下面是一个声明请求数据是一个`UserReq`其结构中包含`username`、`passwd`、`age`三个字段。这三个字段类型分别是`string`、`string`、`int64`。

```api
type UserReq{
    username string `form:"username"`
    passwd string   `form:"passwd"`
    age int64   `form:"form"`
}
```

值得注意的是，类型后面的叫做标签(golang中的TAG,TAG的具体用法请自行查阅相关资料)，标签中`form`为这个数据将会从请求的Url参数中获取，`json`标签表示数据将会从请求体中获取。
当然也可以使用如下语法，同时生成多个结构体:

```api
type(
    UserReq1{
        username string `form:"username"`
        passwd string   `form:"passwd"`
        age int64   `form:"form"`
    }
    UserReq2{
        username string `form:"username"`
        passwd string   `form:"passwd"`
        age int64   `form:"form"`
    }
    UserReq3{
        username string `form:"username"`
        passwd string   `form:"passwd"`
        age int64   `form:"form"`
    }
)
```

5. 服务声明: 在Gozero中，多个API方法组成一个服务，我们可以通过下面的方法声明一个服务:

```api
@server(
    jwt:   Auth // 代表当前服务均会开启jwt鉴权
    group: foo // 当前服务的分组名称
)
```
下面我们继续给这个服务中添加具体API方法:

```api
@server(
    jwt:   Auth // 代表当前服务均会开启jwt鉴权
    group: foo // 当前服务的分组名称
)
// 一个service分组
service user-api{
    // 接口描述
    @doc "foo"
    // 接口处理方法名称
    @handler foo
    // 接口: 请求方法、路由、请求入参、返回值
    post /foo (Foo) returns (Bar)
}
```
>上面的示例中: 我们可以看到在接口声明时候，结构如下为`请求方法 请求路径 请求参数 返回值` 这里的返回值和请求参数均为我们上面声明的数据结构的名称。请求方法分为常见的Restful请求方法，如(post、delete、get、put等)


#### API文件生成结构
由于框架提供了一套定义API文件的基本语法，因此直接使用框架工具进行生成相对应的基本桩代码: `goctl api go -api xxxx.api -dir . -style gozero`<br>
生成之后的目录文件如下:

```txt
.
├── etc
│   └── foo-api.yaml // 生成的配置文件
├── example.api // 你的api文件
├── foo.go // 主程序main
├── go.mod // golang 模块声明
└── internal // 内部包，对外部模块不可见
    ├── config // 配置文件文件夹
    │   └── config.go //配置文件对应的go结构
    ├── handler // 处理器
    │   ├── foo // 声明的Server分组
    │   │   └── foohandler.go 
    │   └── routes.go
    ├── logic // 实际业务逻辑
    │   └── foo
    │       └── foologic.go
    ├── svc // 上下文
    │   └── servicecontext.go
    └── types // 声明的API文件中结构对应的Go结构体
        └── types.go
```
由于存在API文件的定义，框架会自动帮助我们封装相应的返回参数，因此只需要编写相对应的logic即可
### 编程要求
请按照下面要求，编写对应的API文件，并且使用Goctl工具生成对应的API项目结构。
1. API 文件名称为`user.api`
2. 文件中需要定义如下两个实体: 标签中名称与属性名称一致
    - `UserLoginReq`
        - 用户名: `userName` 类型 string
        - 密码: `passwd` 类型 string
    - `UserLoginResp` 
        - Token: `token` 类型string
3. 定义一个名称为`user`的服务组 && 名称为`user-api`的服务(不要开启Jwt认证)
4. 在`user-api`服务中包含一个API方法，,接口描述为:`doc`,方法名称为`Login` 请求地址为`/login` 入参和出参为上述结构实体。
5. 编写完成API文件之后，请在当前目录使用`goctl`工具生成对应的api结构,先进入到编辑的文件目录: `step/step3/stu`之后再生成对应的结构。
6. 修改生成文件目录下: step/step3/example/etc/user-api.yaml 配置文件中端口号为8889
7. 修改:`step/step3/example/internal/logic/user/loginlogic.go` 编写相应逻辑，使其最终通过API返回的数据中token为：我发送的userName + passwd
---
### 测试说明
平台会对你编写的代码进行测试，通过向你的`8889`端口发送上述请求示例的请求，如果得到的响应与上述响应数据一致，则测评通过，否则失败
#### 示例
请求参数 http://localhost:8889/login
post:
```json
{"userName":"abcd","passwd":"1234567"}
```
响应数据:
```json
{"token":"abcd1234567"}
```
---
开始你的任务吧，祝你成功！


