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
