/*
@Time : 2020-01-28 12:01 
@Author : Tenlu
@File : api
@Software: GoLand
*/
package routers

import (
	"gin-learn-todo/app/controller/user"
	"gin-learn-todo/app/middleware"
	"github.com/gin-gonic/gin"
)


func api(r *gin.Engine) {
	// middleware.RabcMiddleware()
	api := r.Group("/api", middleware.Cors())
	{
		api.GET("/user/info", user.Info)
		api.GET("/user/list", user.List)
	}
}
