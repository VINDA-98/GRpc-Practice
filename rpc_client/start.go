package rpc_client

import (
	"log"
	"net/rpc"
)

// @Title  rpc_server
// @Description  MyGO
// @Author  WeiDa  2023/5/11 16:34
// @Update  WeiDa  2023/5/11 16:34

type RpcClient struct {
	Client *rpc.Client
}

func (c *RpcClient) StartClient() (err error) {
	c.Client, err = rpc.DialHTTP("tcp", "127.0.0.1:8088")
	if err != nil {
		return
	}

	if err != nil {
		log.Fatalln("dialing error:", err)
	}
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
