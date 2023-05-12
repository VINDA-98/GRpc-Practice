package main

import (
	"github.com/VINDA-98/gRpc_study/rpc_client"
	"github.com/VINDA-98/gRpc_study/rpc_server"
	"log"
	"time"
)

// @Title  gRpc_study
// @Description  MyGO
// @Author  WeiDa  2023/5/10 18:43
// @Update  WeiDa  2023/5/10 18:43

func main() {
	practice1()
	//启动rpc服务端
	serverRPC := &rpc_server.RpcServer{}
	serverGRPC := &rpc_server.RpcServer{}
	quitMain := make(chan bool) //阻塞main goroutine,等待for select 处理完成

	go func() {
		_ = serverRPC.StartServer()
	}()
	go func() {
		_ = serverGRPC.StartGRPCServer()
	}()

	time.Sleep(time.Second * 2)
	//启动rpc客户端
	clientRPC := &rpc_client.RpcClient{}
	clientGRPC := &rpc_client.RpcClient{}
	_ = clientRPC.StartClient()
	_ = clientGRPC.StartGRPCClient()

	// 定义一个定时器，3S后自动请求rpc服务端
	ticker1 := time.NewTicker(time.Second * 2)

	go func() {
		for {
			select {
			//启动3S后，开始请求
			case <-ticker1.C:
				ticker1.Stop()
				for i := 0; i < 3; i++ {
					err := clientRPC.SendClient()
					if err != nil {
						log.Println("SendClient:", err)
					}
					err = clientGRPC.SendGRPCClient()
					if err != nil {
						log.Println("SendGRPCClient:", err)
					}
				}
				_ = clientRPC.CloseClient()
				_ = clientGRPC.CloseGRPCClient()
				_ = serverRPC.CloseServer()
				_ = serverGRPC.CloseGRPCServer()
				quitMain <- true
			}
		}
	}()

	<-quitMain

}
