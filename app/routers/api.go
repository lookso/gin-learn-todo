/*
@Time : 2020-01-28 12:01 
@Author : Tenlu
@File : api
@Software: GoLand
*/
package routers

import (
	"gin-learn-todo/app/controller"
	"github.com/gin-gonic/gin"
)

type api struct{}

func (_ api) Load(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/post/info", controller.Info)
		api.GET("/post/list", controller.List)
	}
}
