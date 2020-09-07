/*
@Time : 2019-11-24 11:00
@Author : Peanut
@File : server
@Software: GoLand
*/
package boot

import (
	"context"
	"fmt"
	"gin-learn-todo/routers"
	"gin-learn-todo/setting"
	"github.com/getsentry/sentry-go"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func NewServer() {
	initSentry()
	initGin()
	initPrometheus()
}

func initGin() {
	router := gin.New()
	routers.All(router)
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("logs/gin.log")
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

	router.Use(gin.Recovery(), sentryGin.New(sentryGin.Options{Repanic: true}))

	// 先开启release mode ，屏蔽掉gin默认的waring
	gin.SetMode(gin.ReleaseMode)
	if setting.Conf.App.Debug {
		pprof.Register(router) // 性能分析工具
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化http server
	s := &http.Server{
		Addr:              setting.Conf.App.Addr,
		Handler:           router,
		ReadTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 60,
		MaxHeaderBytes:    1 << 20, //1M
	}
	wg := sync.WaitGroup{}

	// 启动http服务
	go func() {
		wg.Add(1)
		defer wg.Done()
		// service connections
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("http.server start success,listen.port%s", setting.Conf.App.Addr)
	if setting.Conf.App.Debug {
		// debug 模式下开启gops
		log.Println("gops listen at ", ":9000")
		if err := agent.Listen(agent.Options{
			Addr:            ":9000",
			ShutdownCleanup: true, // automatically closes on os.Interrupt
		}); err != nil {
			log.Println("gops agent %s\n", err)
		}
	}
	Quit(s)
	wg.Wait()
}

func initPrometheus() {
	p := NewPrometheus("gin")
	p.Use(&gin.Engine{})
}

func Quit(s *http.Server) {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown error: ", err)
	}
	log.Println("server exiting success")
}

func initSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              setting.Conf.Sentry.Dsn,
		AttachStacktrace: true,
	}); err != nil {
		log.Fatalf("sentry initialization failed: %v", err)
	}
	log.Println("init sentry success")
}
