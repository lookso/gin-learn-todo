/*
@Time : 2020-02-01 14:04 
@Author : peanut
@File : rabc
@Software: GoLand
*/
package middleware

import (
	"github.com/gin-gonic/gin"
)

func RabcMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		//	email := "peanut@itech8.com"
		//	appId := setting.Conf.ApiServer.AppId
		//	var app model.App
		//	info, err := app.GetAppById(appId)
		//	if info == nil || err != nil {
		//		c.AbortWithStatusJSON(response.BadRequest("应用Id错误"))
		//		c.Abort()
		//		return
		//	}
		//	b, err := model.Verifier(appId, email, c.Request.URL.Path, c.Request.Method)
		//
		//	if b != true || err != nil {
		//		c.AbortWithStatusJSON(response.Forbidden("forbidden"))
		//		c.Abort()
		//		return
		//	}
		c.Next()
	}
}
