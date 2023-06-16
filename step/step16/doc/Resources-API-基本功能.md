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

### 难点介绍
#### 递归
其是一种算法或过程，在该算法中，函数重复调用自身以解决问题的方法。递归函数通常解决问题的步骤也是逐步递推的，每一次递归调用都会让问题规模变小直到结束条件满足，递归函数才会返回结果。
下面我们通过示例，演示如何使用递归计算给定数字的阶乘。
```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```
函数接收一个整数作为参数，并检查这个整数是否等于0。如果相等，则返回1；否则，它会递归地调用自己，计算n-1的阶乘并将其与n相乘，最后返回结果。
如果我们调用`factorial(5)`，该函数将递归调用自己5次，每次调用都将返回以下结果：
```
factorial(5) = 5 * factorial(4) = 5 * 4 * factorial(3) = 5 * 4 * 3 * factorial(2) = 5 * 4 * 3 * 2 * factorial(1) = 5 * 4 * 3 * 2 * 1 * factorial(0) = 5 * 4 * 3 * 2 * 1 * 1 = 120
```
因此，函数的返回值将是120，这是5的阶乘。
#### 协程
> 更多参考:
> 1. [Go语言协程使用最佳实践](https://zhuanlan.zhihu.com/p/374464199)
> 2. [Channel](https://www.runoob.com/w3cnote/go-channel-intro.html)
> 3. [select](https://www.topgoer.com/%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6/%E6%9D%A1%E4%BB%B6%E8%AF%AD%E5%8F%A5select.html)
##### Goroutine
Go 协程（Goroutine）是与其他函数同时运行的函数。又称轻量级线程，是一种用户态的轻量级线程，具有独立的状态和上下文切换。协程的优势在于它能够避免操作系统线程切换的开销，提升程序并发性能。<br>
与线程不同的是，协程是由程序员控制的，通过协作式调度实现多任务并发，而线程则是由操作系统调度的，通过时间片轮转实现多任务并发。 协程与进程和线程的最大不同在于它是由程序员控制的，程序员可以随时将控制权交给其他协程，从而实现协作式多任务并发。<br>
另外，协程可以在单线程中执行多个任务，避免了多线程并发带来的竞争和锁问题，同时也不会造成过多的资源占用和调度开销。而进程和线程则需要占用独立的地址空间和系统资源，系统对其进行调度和管理。
下面给出一个协程的基本使用示例:
```go
func main() {
    go func() {
        fmt.Println("goroutine")
    }()
    fmt.Println("main")
    time.Sleep(time.Second)
}
```
上述代码中，我们使用关键字go开启一个协程，输出“goroutine”并立即返回，然后主协程继续执行，输出“main”，最后通过Sleep函数等待1秒钟，使得所有的协程都有足够的时间执行完毕。
通过对以上示例的分析，我们可以看出，在golang中使用协程非常方便，可以轻松地实现并发编程。同时，与传统的线程相比，协程具有更高的效率和更好的并发性能，是现代并发程序开发中不可或缺的一部分。
##### Channel
其被认为是协程之间通信的管道。与水流从管道的一端流向另一端一样，数据可以从信道的一端发送并在另一端接收。 每个 channel 都有一个类型。此类型是允许信道传输的数据类型。<br>
channel 是类型相关的，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。 其也是种基于通信的同步机制，在并发编程中起着重要的作用。<br>
channel的本质是一个队列，可以实现先进先出的数据传递方式，可以实现同步和异步两种模式。channel可以实现数据的同步发送和接收，防止数据的竞争和冲突。
channel的基本使用方法：
1、定义channel变量： `var c chan int`
2、创建`channel：c = make(chan int)`，也可以使用简化形式，`c := make(chan int)`
3、发送数据：`c <- 1`
4、接收数据：`x := <- c` 从channel c接收一个值，并赋值给变量`x`
下面给出一个基本使用示例:
```go
func main() {
    c := make(chan int)
    go func() {
        fmt.Println("Hello world")
        c <- 1
    }()
    <-c // 等待channel有数据
    fmt.Println("Done")
}
```
##### Select
在golang中，使用关键字select可以实现多个协程的运行和协调，它可以监听多个channel的读写操作，并在其中一个channel可以操作时立即执行该操作。
下面给出使用select实现多个协程运行的示例：
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 1
    }()

    go func() {
        ch2 <- 2
    }()

    select {
    case <-ch1:
        fmt.Println("channel 1 received")
    case <-ch2:
        fmt.Println("channel 2 received")
    }
}
```
在上述示例中，我们创建了两个channel，分别用于两个协程进行通信。第一个协程将数字1发送到ch1中，第二个协程将数字2发送到ch2中。<br>
之后通过使用select语句监听这两个channel的读操作，并在其中一个channel可以读取时立即执行相关操作。在本例中，我们使用了两个case语句来监听ch1和ch2的读操作，只有其中一个channel可以读取时，就会执行相关case语句中的操作，并输出相应信息。
通过使用select语句，我们可以方便地实现多个channel的并发读写操作。当有多个协程需要协调时，它可以用作一种非常有效的调度工具，提高程序的并发性能。<br>
但是如果所有的case语句都不能进行读写操作，那么select语句会阻塞程序，等待其中一个case语句可以进行读写操作时才会继续执行下去。在上面的示例中，有两个协程都在发送数据到对应的channel中，因此其中一个case语句一定可以进行读取操作，不会发生阻塞的情况。
如果我们需要在select语句中对超时等操作进行处理，就需要借助于Golang中常用的time包中的函数，例如time.After函数可以实现在指定时间后返回一个channel，从而可以在select语句中使用。
下面给出基本使用示例:
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- 1
    }()

    select {
    case <-ch1:
        fmt.Println("channel 1 received")
    case <-ch2:
        fmt.Println("channel 2 received")
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
    }
}
```
在以上示例中，我们创建了一个在2秒后向ch1发送数据的协程，并在select语句中监听了ch1和ch2的读操作。但由于ch1的读取操作需要等待2秒钟，而由于我们在select语句中设置了一个超时时间为1秒的case语句，所以当select语句等待的时间超过1秒时，会执行超时的语句，并输出"timeout"。
因此，我们可以通过在select语句中加入time.After函数的方式实现超时的处理，从而避免程序在长时间等待时发生阻塞。


### 编程要求
#### 提示
1. 使用到的工具类:
   - `sources/common/utils/oss_utils.go`
   - `sources/common/utils/file_utils.go`
2. 树形分类`GetClassificationAlL`
在本关卡内，我们需要查询响应的树形分类，其每个树形分类都可能具有子节点，因此需要通过递归的方式进行查询。具体递归函数的示例如上。
3. 文件上传解析，相对示例请参考前一关卡。
4. 请修改对应的配置文件，配置对应的Rpc连接地址(如是按照前面关卡要求完成的编写，只需要检查此部分`sources/apps/resources/cmd/api/etc/resources.yaml`配置是否与下述一致即可):
```yaml
ResourcesRpc:
   EndPoints:
      - 0.0.0.0:8011
OrderRpc:
   Endpoints:
      - 0.0.0.0:8010
TradeRpc:
   Endpoints:
      - 0.0.0.0:8012
```
5. 配置自己的阿里云OSS对应的密钥
```yaml
AliCloud:
  AccessKeyId: 
  AccessKeySecret: 
  BucketName: 
  EndPoint: 
  CommonPath: resources/files
  CachePath: cache/poster
```
#### 任务
完成logic部分全部函数功能编写，具体如下:
````text
.
├── classification // 分类逻辑
│   ├── getallclassificationslogic.go // 查询所有分类数据(树形菜单)
│   └── getclassificationdatabypageslogic.go // 通过分页查询所有分类数据
├── comment
│   ├── addusercommentlogic.go // 增加用户评论
│   ├── deleteusercommentlogic.go // 删除用户评论
│   ├── getcommentdetailbyidlogic.go // 查询用户评论详细信息
│   └── getcommentsconditionlogic.go // 条件查询用户评论
├── deplete
│   └── getfileandspaceinsightlogic.go // 查询用户空间使用情况
├── fileRes
│   ├── deletefilereslogic.go // 删除用户文件
│   ├── getfileresdatalogic.go // 获取用户文件数据（实际文件，需要判断用户是否购买了该款文件，请调用前面实现的OrderRPC）
│   ├── getfileresinfologic.go // 查询用户文件详细信息
│   ├── searchfilereslogic.go // 条件搜索文件
│   └── uploadfilereslogic.go // 上传文件信息，文件会上传到OSS，并存储相关信息到db
└── textRes
    ├── deletetextreslogic.go // 删除文本资源
    ├── gettextinfologic.go // 获取文本资源
    ├── searchtextreslogic.go // 查询文本资源
    └── uploadtextreslogic.go // 是上传文本资源，包括文本资源的头像数据
````
### 测试说明
测试前提: 平台会编译前面任务完成的`user-rpc`,`user-api`,`order-rpc`,`trade-rpc`,`resource-rpc`,`resources-api`服务，之后启动当前关卡的服务进行测评<br>
请在完成**所有接口**之后，在进行评测，平台会评测上述要求中的**部分接口**，**其结果可以满足平台进行增删改查配置的标准输入输出时，视为通过测评，否则不给予通过**
会被评测的接口内容如下:
- `/user/v1/login`
- `/res/v1/file/uopload`
- `/res/v1/classification/getAll`
- `/res/v1/classification/subDatas`