package response

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int32       `json:"total,omitempty"`
	Data  interface{} `json:"data"`
}

func DataWithTotal(total int32, obj interface{}) (int, Response) {
	resp := Response{
		Code:  0,
		Msg:   "",
		Data:  obj,
		Total: total,
	}
	return 200, resp
}

func Data(obj interface{})(int,Response) {
	resp:=Response{
		Code:0,
		Msg:"",
		Data:obj,
	}
	return 200,resp
}

func Error(code int, msg string) (int, Response) {
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	return 200, resp
}