/*
@Time : 2020-01-28 12:01
@Author : Tenlu
@File : api
@Software: GoLand
*/
package routers

import (
	"gin-learn-todo/controller"
	"gin-learn-todo/middleware"
	"github.com/gin-gonic/gin"
)

func api(r *gin.Engine) {
	l := r.Group("/api/v1", middleware.Cors())
	{
		l.POST("/user/login", controller.Login)
	}

	api := r.Group("/api/v1", middleware.JwtAuth())
	{
		api.GET("/user/info", controller.Info)
		api.GET("/user/list", controller.List)
	}
}
