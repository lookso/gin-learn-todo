/*
@Time : 2019-11-17 10:18 
@Author : Tenlu
@File : api
@Software: GoLand
*/
package router

import (
	"gin-learn-todo/app/controller"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (_ Api) load(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user/list", controller.FindUser)
		api.GET("/user/edit", controller.UpdateUser)
		api.POST("/user/batchinsert", controller.BatchInsertUser)
		api.POST("/user/insert", controller.InsterUser)
	}
}

