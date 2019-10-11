package router

import (
	"fmt"
	"gin-sourcecode-learn/controller"
	"gin-sourcecode-learn/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	router.RedirectFixedPath=true
	fmt.Println("base path:",router.BasePath())

	router.LoadHTMLGlob("./template/*") // html模板
	// 中间件
	router.Use(middleware.Login("admin"), gin.Recovery())

	router.GET("ping", controller.GetPing)
	router.GET("view", controller.GetHtml)

	userGroup := router.Group("/user")

	userGroup.Handle("GET", "/test", controller.UserInfo,controller.GetPing)

	userGroup.POST("/create", controller.InsertUser)
	userGroup.GET("/detail", controller.UserInfo)
	userGroup.GET("/JSONP?callback=x", controller.GetJsonp)

	fmt.Println("userGroup:",userGroup.BasePath())

	// url重定向
	router.GET("/redirect", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// 路由组使用 gin.BasicAuth() 中间件
	// 模拟一些私人数据
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",    //  用户:密码
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
