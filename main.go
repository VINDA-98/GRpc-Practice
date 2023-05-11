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
	server := &rpc_server.RpcServer{}
	quitMain := make(chan bool) //阻塞main goroutine,等待for select 处理完成

	go server.StartServer()

	time.Sleep(time.Second * 3)
	//启动rpc客户端
	client := &rpc_client.RpcClient{}
	_ = client.StartClient()

	// 定义一个定时器，3S后自动请求rpc服务端
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			select {
			//启动3S后，开始请求
			case <-ticker.C:
				for i := 0; i < 3; i++ {
					err := client.SendClient()
					if err != nil {
						log.Println("SendClient:", err)
					}
				}
				_ = client.CloseClient()
				server.CloseServer()
				quitMain <- true
			}
		}
	}()

	<-quitMain

}
