package routers

import (
	"gin-learn-todo/app/setting"
	"github.com/getsentry/sentry-go/gin"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {

	router := gin.New()

	// 默认中间件注册
	router.Use(gin.Recovery(), sentryGin.New(sentrygin.Options{Repanic: true}))

	// 先开启release mode ，屏蔽掉gin默认的waring
	gin.SetMode(gin.ReleaseMode)
	if setting.Conf.ApiServer.Debug {
		pprof.Register(router) // 性能分析工具
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
		return
	})

	return router
}
func initRouter(r *gin.Engine) {
	api{}.Load(r)
	docs{}.Load(r)
}
