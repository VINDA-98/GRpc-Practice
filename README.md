# gRpc_study
gRPC 学习 ，安装命令如下：
# go编译工具 最新版
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

将proto文件编译为任何语言的文件

# 指定版本
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.3.0

可以使用下面的命令来生成代码
$ protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto


# ProtoBuf 安装
V3 版本： `https://github.com/protocolbuffers/ProtoBuf/releases`

Mac快速安装：
`brew install ProtoBuf`


# proto文件初始化
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



# 编译命令解析
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

## 运行程序
```shell
go build -o gRpcTest
./gRpcTest
```

或者

```shell
chmod a+x start.sh
bash start.sh
```