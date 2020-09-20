package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the router",
		})
		return
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the method",
		})
		return
	})
	//r.Use(middleware.ZipKin())
	// 注册路由
	api(r)
	test(r)
	// 文档
	docs(r)

}
