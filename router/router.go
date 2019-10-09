package router

import (
	"gin-sourcecode-learn/controller"
	"gin-sourcecode-learn/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 模拟一些私人数据
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	router := gin.New()
	router.LoadHTMLGlob("./template/*") // html模板
	// 中间件
	router.Use(middleware.Login("admin"), gin.Recovery())

	router.GET("ping", controller.GetPing)
	router.GET("view", controller.GetHtml)

	userGroup := router.Group("/user")
	userGroup.POST("/create", controller.InsertUser)
	userGroup.GET("/detail", controller.UserInfo)
	userGroup.GET("/JSONP?callback=x", controller.GetJsonp)

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	return router
}
