package routers

import (
	"fmt"
	"gin-learn-todo/setting"
	"github.com/getsentry/sentry-go/gin"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func SetUpRouter() *gin.Engine {

	router := gin.New()

	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("output/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 默认中间件注册
	// gin.Logger()
	// 自定义日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery(), sentryGin.New(sentrygin.Options{Repanic: true}))

	// 先开启release mode ，屏蔽掉gin默认的waring
	gin.SetMode(gin.ReleaseMode)
	if setting.Conf.ApiServer.Debug {
		pprof.Register(router) // 性能分析工具
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	All(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the router",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not find the method",
		})
		return
	})
	
	return router
}
func All(e *gin.Engine) {
	api(e)
	docs(e)
}
