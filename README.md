# gRpc_study
gRPC 学习
，安装命令如下：
# 最新版
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 指定版本
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.3.0
复制代码

可以使用下面的命令来生成代码
$ protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
