package rpc_client

// @Title  rpc_service
// @Description  MyGO
// @Author  WeiDa  2023/5/11 15:36
// @Update  WeiDa  2023/5/11 15:36

type Arity struct {
}

// ArityRequest 请求结构体
type ArityRequest struct {
	Name     string
	Birthday string //生辰八字
}

// ArityResponse 响应结构体
type ArityResponse struct {
	Level string
	Title string
	Score map[string]int //最后成绩

}
