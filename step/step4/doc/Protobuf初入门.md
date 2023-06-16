# 介绍[TOC]

---

### 任务描述
本关卡任务：编写一个符合Protobuf定义的Protobuf文件

### 相关知识
Protobuf: [官方文档](https://protobuf.dev/)
学习文档: [jitwxs](https://jitwxs.cn/60aca815#2-8-Import)


#### protobuf 基本介绍
Protocol Buffers 是一种轻量级、高效、跨平台的序列化框架，由 Google 公司开发并开源。通过发送数据时将其编码成二进制格式，它可以有效地减小数据传输体积，同时也可以保证数据在传输过程中的完整性。其数据结构是通过一个结构化的描述文件来定义的，这个文件使用 protobuf 的语言来编写，并描述了要序列化的消息的结构、字段类型和标记信息。在序列化数据时，程序会将数据按照描述文件中的定义编码成二进制格式，这些数据可以通过网络或存储在文件中进行传输和存储。
优点：
- 包括支持语言丰富
- 序列化速度快
- 生成的代码短小、易于维护和扩展等

##### protobuf 数据类型
|Protobuf 数据类型 |描述|	Golang 数据类型|
|---|---|---|
|double| |	float64|
|float	||float32|
|int32	|使用可变长度编码。对负数编码效率低下|int32|
|int64	|使用可变长度编码。对负数编码效率低下|int64|
|uint32	|使用可变长度编码。|uint32|
|uint64	|使用可变长度编码。|uint64|
|sint32	|使用可变长度编码。有符号的 int 型|int32|
|sint64	|使用可变长度编码。有符号的 int 型|int64|
|fixed32|四个字节	|uint32|
|fixed64|八个字节	|uint64|
|sfixed32|总是四个字节|	int32|
|sfixed64|总是八个字节	|int64|
|bool	||bool|
|string	|UTF-8 编码或 7 位 ASCII 文本，不长于 232|string|
|bytes	|任意字节 不长于 232|[]byte|
##### protobuf 基本语法
每一个protobuf文件中都，包含一下几个部分:
- 文件语法声明`syntac`: 表名文件使用的potobuf版本`syntax = "proto3";`
- 包名称声明`package`: 防止Message之间的定义名称冲突
- 导入`import`: 可以通过import导入其他包的文件 `import "xxxx/other_protos.proto";`
- 对于各语言的包声明: 如`java_package`、`go_package`等
- 消息声明`message`: 表示这个是一个protobuf文件中的结构化对象 
- rpc方法声明: 声明一个rpc调用方法，入参与出参均为，声明的message

完整示例:
```proto
syntax = "proto3";

option go_package = "./pb";

package pb;

import "path/to/custom.proto";

enum Code {
  SUCCESS = 0;
  FAILURE = 1;
}

message User {
  string name = 1;
  int32 age = 2;
}

message Order {
  string order_id = 1;
  repeated string item_ids = 2;
}

service OrderService {
  rpc GetOrdersForUser(User) returns (Order);
  rpc GetAllOrdersForUser(User) returns (stream Order);
}

```
###### 枚举
各种语言之中，都可以通过声明枚举值的形式来进行存储变量，在portobuf里面也可以实现相同的定义，下面通过一个例子快速了解
> Protobuf 中，枚举类型的值必须唯一，且不可以使用保留的数字（如 -1）。

```proto
syntax = "proto3";
package example;

message TestMessage {
  enum Color {
    RED = 0;
    GREEN = 1;
    BLUE = 2;
  }
  
  Color color = 1;
}

```

##### 一般消息
消息声明使用messgae + 结构名称进行声明，其中消息中字段声明必须为（类型 名称），字段后面需要加入消息序号，用来在消息的二进制格式中识别各个字段的，一旦开始使用就不能够再改变。[1,15]之内的标识号在编码的时候会占用一个字节。[16,2047]之内的标识号则占用2个字节。
示例:
```proto
message Test{
    string Name = 1;
    int64 Id = 2;
}
```

##### 消息嵌套
消息与消息之间，可以类似面向对象语言中的嵌套
示例：
```proto
message Test{
    string Name = 1;
    int64 Id = 2;
    TestInner inner = 3;
}

message TestInner{
    string name = 1;
}
```
在这个示例中，我们将TestInner嵌套在Test内部

##### 列表
消息可以按照如下定义实现，属性成为数组或者列表
示例：
```proto
message Test{
    string Name = 1;
    int64 Id = 2;
    repeated TestInner inner = 3;
}
```

##### Map类型
声明方法如下: 这里不能使用repeated进行修饰，枚举也不能是K。

```proto
map<K, V> map_field = N;
```
示例:
```proto
message Test{
    string Name = 1;
    int64 Id = 2;
    map<string,int64> name_id_Map = 3; // 昵称和id的Map
}
```

##### 服务定义
使用 RPC（远程过程调用），文件中是可以定义 RPC 服务接口，生成的RPC调用代码，将和最后指定生成的编程语言相关。
示例:
```proto
syntax = "proto3";

option go_package = "./pb";

package pb;

message Test{
    string Name = 1;
    int64 Id = 2;
    map<string,int64> name_id_Map = 3; // 昵称和id的Map
}
message User {
    string name = 1;
    int64 uid = 2;
}

// 定义User服务
service UserService {
  rpc GetUserFromTest(Test) returns (User);
}
```
在上面的示例中，我们编写了一个UserService的Rpc服务，其中包含一个RPC方法，叫做GetUserFromTest，其方法入参是Test，返回值是User。



#### 生成proto命令使用
> 本实验环境已经预装protobuf ,如需自己部署请参杂下面链接，仅给出golang版本，其余版本请自行google
> 安装教程: https://juejin.cn/post/7101304169524527140
```sh
protoc --go_out=. xxx.proto   //生成Go代码
protoc --java_out=. xxx.proto   //生成Java代码
protoc --cpp_out=. xxx.proto   //生成C++代码
```
之后将会生成各个语言下的protobuf文件包，如go语言的`xxx.pb.go`

#### Go语言的protobuf文件使用
示例:
1. 编写protobuf:
```proto
syntax = "proto3";
package message;

message Person {
  string name = 1;
  int32 age = 2;
  repeated string hobbies = 3;
}
```
2. 编译proto: `protoc --go_out=. person.proto `
3. 编写golang使用:
```go
//Go语言示例
package main

import (
    "fmt"
    "github.com/golang/protobuf/proto"
    "./person.pb"
)

func main() {
    person := &person.Person{
        Name: "Alice",
        Age:  18,
        Hobbies: []string{"reading", "music", "sports"},
    }

    data, _ := proto.Marshal(person) // protobuf序列化
    newPerson := &person.Person{}
    proto.Unmarshal(data, newPerson)  // protobuf反序列化
    fmt.Println(newPerson.GetName(), newPerson.GetAge(), newPerson.GetHobbies())
}

```
### 编程要求
经过上面的学习，我们已经基本了解了Protobuf的基本语法，下面请按照要求编写一个Probuf文件，并生成相应的go 桩代码。
> 注意：编辑的测试文件是: step/step4/stu/user.proto，之后再前目录执行，请不要修改测试文件名称，否则会影响评测！！！
1. Proto文件中包含两个Message
- UserReq:
    - UserName(string)
    - UserId(int64)两个参数
- UserResp:
    - UserHome(string)
    - UserMoneyCardNum(int64): 注意这个UserMoneyCardNum存在多个。
2. 定义一个RPC方法，方法签名是，`GetUserInfo`入参是UserReq，返回值是UserResp。

### 测试说明
平台会测试比对： Proto文件的具体内容和标准答案是否一致（去除空格）