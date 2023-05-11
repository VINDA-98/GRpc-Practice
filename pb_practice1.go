package main

import (
	"fmt"
	"github.com/VINDA-98/gRpc_study/protocol"
	"google.golang.org/protobuf/proto"
)

// @Title  pb
// @Description  MyGO
// @Author  WeiDa  2023/5/11 12:22
// @Update  WeiDa  2023/5/11 12:22

// practice1 练习1
func practice1() {
	scopeMap := make(map[string]int32, 3)
	scopeMap["Golang"] = 100
	scopeMap["Java"] = 100
	scopeMap["C++"] = 100
	userInfo := &protocol.UserInfo{
		Name:     "VinDa",
		Id:       2963,
		Sex:      1,
		Duration: 7200,
		Score:    scopeMap,
	}
	//用字符串的方式：打印ProtoBuf消息
	fmt.Printf("字符串输出结果：%v\n", userInfo.String())

	//转成二进制文件
	marshal, err := proto.Marshal(userInfo)
	if err != nil {
		return
	}
	fmt.Printf("Marshal转成二进制结果：%v\n", marshal)

	//将二进制文件转成结构体
	newUserInfo := protocol.UserInfo{}
	err = proto.Unmarshal(marshal, &newUserInfo)
	if err != nil {
		return
	}
	fmt.Printf("二进制转成结构体的结果：%v\n", &newUserInfo)
}
