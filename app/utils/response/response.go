package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(obj interface{}) Response {
	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: obj,
	}
	return resp
}

func Error(code int, msg string) Response {
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	return resp
}
