package rpc_server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// @Title  rpc_server
// @Description  MyGO
// @Author  WeiDa  2023/5/11 16:34
// @Update  WeiDa  2023/5/11 16:34

type RpcServer struct {
	Listener net.Listener
}

func (s *RpcServer) StartServer() {

	//注册rpc服务
	err := rpc.Register(new(Arity))
	if err != nil {
		log.Fatalln("Register error:", err)
	}

	rpc.HandleHTTP() //采用http协议作为rpc载体

	s.Listener, err = net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatalln("Listen error:", err)
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s", "正在新建RPC服务...\n")

	//常规启动http服务
	err = http.Serve(s.Listener, nil)
	if err != nil {
		return
	}

}

func (s *RpcServer) CloseServer() {
	//关闭rpc服务
	err := s.Listener.Close()
	if err != nil {
		log.Fatalln("CloseServer error:", err)
	}

	log.Println("正在关闭RPC服务...")
}
