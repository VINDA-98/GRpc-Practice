package test

import (
	"github.com/VINDA-98/gRpc_study/rpc_client"
	"github.com/VINDA-98/gRpc_study/rpc_server"
	"log"
	"testing"
)

func TestRpcServer(t *testing.T) {
	server := &rpc_server.RpcServer{}
	server.StartServer()
}

func TestGRpcServer(t *testing.T) {
	server := &rpc_server.RpcServer{}
	err := server.StartGRPCServer()
	if err != nil {
		return
	}
}

func TestRpcClient(t *testing.T) {
	//启动rpc客户端
	client := &rpc_client.RpcClient{}
	_ = client.StartClient()
	for i := 0; i < 3; i++ {
		err := client.SendClient()
		if err != nil {
			log.Println("SendClient:", err)
		}
	}
	_ = client.CloseClient()

}
