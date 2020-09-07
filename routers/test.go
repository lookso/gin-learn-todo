package routers

import (
	"gin-learn-todo/controller"
	"github.com/gin-gonic/gin"
)

func test(r *gin.Engine) {
	// 获取原始数据流
	r.GET("/api/v1/req", controller.GetRawData)
}
