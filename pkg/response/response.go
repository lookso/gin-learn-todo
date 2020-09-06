package response

import "net/http"

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int32       `json:"total,omitempty"`
	Data  interface{} `json:"data"`
}

func DataWithTotal(Total int32, obj interface{}) (int, Response) {
	resp := Response{
		Code:  0,
		Msg:   "success",
		Total: Total,
		Data:  obj,
	}
	return 200, resp
}
func Data(obj interface{}) (int, Response) {
	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: obj,
	}
	return 200, resp
}

func Error(code int, msg string) (int, Response) {
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	return 200, resp
}

// ParamsError 参数错误，报错参数格式错误，参数值异常等导致数据找不到
func ParamsError(msg string) (int, Response) {
	if msg == "" {
		msg = "参数有误"
	}
	resp := Response{
		Code: ErrCodeParamsError,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusOK, resp
}

// BadRequest 请求不完整，缺少必要参数等
func BadRequest(msg string) (int, Response) {
	if msg == "" {
		msg = "缺少参数"
	}
	resp := Response{
		Code: ErrCodeBadRequest,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusBadRequest, resp
}

// UnAuthorized 返回
func UnAuthorized(msg string, data interface{}) (int, Response) {

	if msg == "" {
		msg = "请先登录"
	}

	resp := Response{
		Code: ErrCodeUnauthorized,
		Msg:  msg,
		Data: data,
	}
	return http.StatusUnauthorized, resp
}

// Forbidden 无权限访问
func Forbidden(msg string) (int, Response) {

	if msg == "" {
		msg = "无权访问"
	}

	resp := Response{
		Code: ErrCodeForbidden,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusForbidden, resp
}

// ServerError 服务器异常，比如调用依赖接口失败，无法正常返回数据
func ServerError(msg string) (int, Response) {

	if msg == "" {
		msg = "访问错误"
	}
	resp := Response{
		Code: ErrCodeServerError,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusInternalServerError, resp
}
