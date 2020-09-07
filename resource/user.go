package resource

// gin 使用的内置验证器:https://mp.weixin.qq.com/s/6-cYukhOavFr-oCn0m3ksA
type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfoResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}
