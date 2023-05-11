package rpc_server

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

// CalcBirthday 算命
func (a *Arity) CalcBirthday(req ArityRequest, resp *ArityResponse) error {
	resp.Level = req.Name + ":  绚烂钻石"
	resp.Title = req.Name + "乃是天选之人"
	scope := make(map[string]int, 2)
	scope["幸运指数"] = 10
	scope["倒霉指数"] = 2
	resp.Score = scope
	return nil
}
