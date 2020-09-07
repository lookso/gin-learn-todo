/*
@Time : 2020-01-30 21:30 
@Author : peanut
@File : docs
@Software: GoLand
*/
package routers

import (
	_ "gin-learn-todo/docs"
	"github.com/gin-gonic/gin"
)

var swagHandler gin.HandlerFunc

func docs(r *gin.Engine) {
	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}
}
