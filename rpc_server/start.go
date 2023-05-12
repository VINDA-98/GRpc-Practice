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
