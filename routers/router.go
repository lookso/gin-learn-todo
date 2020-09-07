package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the router",
		})
		return
	})
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the method",
		})
		return
	})
	api(e)
	docs(e)
}
