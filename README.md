# gRpc_study
### go编译工具 最新版
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

将proto文件编译为任何语言的文件

### 指定版本
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.3.0

可以使用下面的命令来生成代码
$ protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto


### ProtoBuf 安装
V3 版本： `https://github.com/protocolbuffers/ProtoBuf/releases`

Mac快速安装：
`brew install ProtoBuf`


# proto项目初始化
[hello.proto](protocol%2Fproto%2Fhello.proto)
```protobuf
syntax = "proto3";  //定义pb版本号

package protocol;      //定义包名

option go_package = "./;protocol";  // 定义go的包名，用于生成.pb.go

message Say{  //消息结构体
  int64           id    = 1;  //int类型
  string          hello = 2;  //string类型
  repeated string word  = 3;  //string[] 数组
  repeated int32 arrays = 4; //int32[] 数组
}

enum SexType //枚举消息类型，使用enum关键词定义,一个性别类型的枚举类型
{
  UnKnow = 0; //proto3版本中，首成员必须为0，成员不应有相同的值
  MALE = 1;   //1男
  FEMALE = 2; //2女  0未知
}

// 定义一个用户消息
message UserInfo
{
  string name = 1; // 姓名字段
  SexType sex = 2; // 性别字段，使用SexType枚举类型
  int64  id       = 3; //id
  int32  duration = 4; //学习的时长 单位秒
  map<string, int32> score = 5;  //学科 分数的map Map 字段是不能使用repeated关键字修饰。
}


// 定义Quote消息
message Quote {
  string url = 1;
  string title = 2;
  repeated string tags = 3; // 字符串数组类型
}

// 定义ListQuote消息
message ListQuote {
  // 引用上面定义的Quote消息类型，作为results字段的类型
  repeated Quote articles = 1; // repeated关键词标记，说明Quotes字段是一个数组
  // 嵌套消息定义
  message Other {
    string url = 1;
  }
  Other others = 2;
}

```

导入其他proto文件

[article.proto](protocol%2Fproto%2Farticle.proto)
```protobuf
syntax = "proto3";

package nesting;

option go_package = "./;protocol";

message Article {
  string url = 1;
  string title = 2;
  repeated string tags = 3; // 字符串数组类型

}
```

[article_list.proto](protocol%2Fproto%2Farticle_list.proto)
```protobuf
syntax = "proto3";
// 导入Article消息定义
import "protocol/proto/article.proto";

package nesting;

option go_package = "./;protocol";

// 定义ListArticle消息
message ListArticle {
  repeated Article articles = 1;
}
```



### 编译命令解析
protoc --go_out=./protocol protocol/proto/*.proto

./protocol 输出pb.go文件的目录相对路径
protocol/proto/*.proto 表示需要被编译的 .proto 文件

```text
  --go_out=OUT_DIR            指定代码生成目录，生成 Go 代码
  --cpp_out=OUT_DIR           指定代码生成目录，生成 C++ 代码
  --csharp_out=OUT_DIR        指定代码生成目录，生成 C# 代码
  --java_out=OUT_DIR          指定代码生成目录，生成 java 代码
  --js_out=OUT_DIR            指定代码生成目录，生成 javascript 代码
  --objc_out=OUT_DIR          指定代码生成目录，生成 Objective C 代码
  --php_out=OUT_DIR           指定代码生成目录，生成 php 代码
  --python_out=OUT_DIR        指定代码生成目录，生成 python 代码
  --ruby_out=OUT_DIR          指定代码生成目录，生成 ruby 代码
```

# RPC 远程过程调用
### 与HTTP对比
RPC主要用于公司内部的服务调用，性能消耗低，传输效率高，服务治理方便。

HTTP主要用于对外的异构环境，浏览器接口调用，APP接口调用，第三方接口调用等。

### 优势
RPC能够跨多种开发工具和平台

RPC能够跨语言调用

RPC能够提高系统的可扩展性，解耦，提高复用

RPC相较于HTTP，传输效率更高，性能消耗更小，自带负载均衡策略，自动实现服务治理

### 添加grpc相关依赖
```shell
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


### 常见错误
`code = DeadlineExceeded desc = context deadline exceeded` 请求超过超时时间

# 服务端处理
[start.go](rpc_server%2Fstart.go)

需要注意以下几个点:

1、结构体RpcServer需要继承protocol.UnimplementedGreeterServer结构体

2、实现hello.proto中定义Greeter服务提供的SayHello()方法,保证入参和出参正确
    
`SayHello(ctx context.Context, in *protocol.HelloRequest) (*protocol.HelloReply, error)`

3、实现GRPC服务的启动和关闭

```go
package rpc_server

import (
	"context"
	"github.com/VINDA-98/gRpc_study/protocol"
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// @Title  rpc_server
// @Description  MyGO
// @Author  WeiDa  2023/5/11 16:34
// @Update  WeiDa  2023/5/11 16:34

type RpcServer struct {
	Listener net.Listener //监听服务
	GRpc     *grpc.Server //GRPC服务
	protocol.UnimplementedGreeterServer
}

func (s *RpcServer) StartServer() (err error) {

	//注册rpc服务
	err = rpc.Register(new(Arity))
	if err != nil {
		log.Fatalln("Register error:", err)
	}

	rpc.HandleHTTP() //采用http协议作为rpc载体

	s.Listener, err = net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln("Listen error:", err)
	}

	log.Printf("正在新建RPC服务...%s\n", s.Listener.Addr())

	//常规启动http服务
	err = http.Serve(s.Listener, nil)
	return
}

func (s *RpcServer) CloseServer() (err error) {
	//关闭rpc服务
	err = s.Listener.Close()
	if err != nil {
		log.Println("CloseServer error:", err)
	}
	return
}

// StartGRPCServer 启动GRPC服务
func (s *RpcServer) StartGRPCServer() (err error) {
	//tcp协议监听指定端口号
	s.Listener, err = net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Fatalf("StartGRPCServer Failed To Listen: %v", err)
		return
	}
	//实例化gRPC服务
	s.GRpc = grpc.NewServer()

	//服务注册
	protocol.RegisterGreeterServer(s.GRpc, s)
	log.Printf("正在新建GRPC服务...%s\n ", s.Listener.Addr())
	//启动服务
	if err := s.GRpc.Serve(s.Listener); err != nil {
		log.Fatalf("StartGRPCServer Failed to Serve: %v", err)
	}

	return
}

// CloseGRPCServer 关闭GRPC服务
func (s *RpcServer) CloseGRPCServer() (err error) {
	err = s.Listener.Close()
	if err != nil {
		return err
	}
	s.GRpc.Stop()
	log.Println("正在关闭GRPC服务...")
	return
}

// SayHello   实现hello接口 GreeterServer
func (s *RpcServer) SayHello(ctx context.Context, in *protocol.HelloRequest) (*protocol.HelloReply, error) {
	return &protocol.HelloReply{Message: "Hello " + in.Name}, nil
}

```


# 客户端
[start.go](rpc_client%2Fstart.go)

需要注意以下几个点:

1、c.GreeterClient = protocol.NewGreeterClient(c.ClientConn) 得到GRPC新的连接

2、c.GreeterClient.SayHello(ctx, &protocol.HelloRequest{Name: "VinDa"}) 发起调用，等待响应

3、实现服务的新建请和关闭请求

```go
package rpc_client

import (
	"context"
	"github.com/VINDA-98/gRpc_study/protocol"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/rpc"
	"time"
)

// @Title  rpc_server
// @Description  MyGO
// @Author  WeiDa  2023/5/11 16:34
// @Update  WeiDa  2023/5/11 16:34

type RpcClient struct {
	Client        *rpc.Client
	ClientConn    *gRPC.ClientConn
	GreeterClient protocol.GreeterClient
}

func (c *RpcClient) StartClient() (err error) {
	c.Client, err = rpc.DialHTTP("tcp", "127.0.0.1:10000")
	if err != nil {
		return
	}

	if err != nil {
		log.Println("dialing error:", err)
		return
	}
	log.Println("正在新建RPC连接...")
	return
}

func (c *RpcClient) SendClient() (err error) {
	req := ArityRequest{"VinDa", "9898"}
	var res ArityResponse

	err = c.Client.Call("Arity.CalcBirthday", req, &res) //运筹帷幄，调用算命方法
	if err != nil {
		log.Fatalln("Call Arity.CalcBirthday error:", err)
		return
	}
	log.Printf("正在验证: %s  %s 的身份信息,返回内容: %s,%s,%v\n", req.Name, req.Birthday, res.Level, res.Title, res.Score)
	return nil
}

func (c *RpcClient) CloseClient() (err error) {
	err = c.Client.Close()
	log.Println("正在断开RPC请求...")
	return
}

// StartGRPCClient 新建GRPC连接
func (c *RpcClient) StartGRPCClient() (err error) {
	//通过gRPC.Dial()方法建立服务连接
	log.Println("正在新建GRPC连接...")
	c.ClientConn, err = gRPC.Dial("127.0.0.1:10001", gRPC.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("StartGRPCClient Not Connect: %v\n", err)
		return
	}
	//实例化客户端连接
	c.GreeterClient = protocol.NewGreeterClient(c.ClientConn)

	return
}

// SendGRPCClient 发送请求信息
func (c *RpcClient) SendGRPCClient() (err error) {
	//设置请求上下文，因为是网络请求，我们需要设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	//客户端调用在proto中定义的SayHello()rpc方法，发起请求，接收服务端响应
	r, err := c.GreeterClient.SayHello(ctx, &protocol.HelloRequest{Name: "VinDa"})
	if err != nil {
		log.Printf("StartGRPCClient Could Not Greet: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return
}

// CloseGRPCClient 关闭GRPC连接
func (c *RpcClient) CloseGRPCClient() (err error) {
	err = c.ClientConn.Close()
	log.Println("正在断开GRPC请求...")
	return
}

```


# 运行程序
```shell
# 注意 --go-grpc_out 参数，生成文件是hello_grpc.pb.go，符合go语言在RPC的文件
protoc -I ./protocol --go-grpc_out=./protocol  ./protocol/proto/*.proto 
go build -o gRpcTest
./gRpcTest
```

或者

```shell
chmod a+x start.sh
bash start.sh
```
