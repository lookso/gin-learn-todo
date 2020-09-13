/*
@Time : 2019-11-24 11:00
@Author : Peanut
@File : server
@Software: GoLand
*/
package boot

import (
	"context"
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/routers"
	"gin-learn-todo/setting"
	"github.com/getsentry/sentry-go"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var engine *gin.Engine

var logger *zap.SugaredLogger

func NewServer() {
	// 初始化logger
	initLogger()

	engine = gin.New()
	routers.All(engine)
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 默认中间件注册
	engine.Use(gin.Logger())
	// 先开启release mode ，屏蔽掉gin默认的waring
	gin.SetMode(gin.ReleaseMode)
	if setting.Conf.App.Debug {
		//pprof.Register(engine) // 性能分析工具 pprof和gops二选一
		// debug 模式下开启gops
		logger.Infof("gops listen at %s", ":9000")
		if err := agent.Listen(agent.Options{
			Addr:            ":9000",
			ShutdownCleanup: true, // automatically closes on os.Interrupt
		}); err != nil {
			logger.Errorf("gops agent %s\n", err)
		}
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化http server
	s := &http.Server{
		Addr:              setting.Conf.App.Addr,
		Handler:           engine,
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
			logger.Errorf("listen: %s\n", err)
		}
	}()
	logger.Infof("http.server start success,listen.port%s", setting.Conf.App.Addr)
	initSentry()
	initPrometheus()

	Quit(s)
	wg.Wait()
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
		logger.Fatal("server Shutdown error: ", err)
	}
	logger.Infof("server exiting success")
}

func initSentry() {
	// 如果没指定sentry开启的环境，则默认 test 和 prod 开启sentry
	if setting.Conf.Sentry.Env == "" {
		setting.Conf.Sentry.Env = "test,prod"
	}
	if setting.Conf.Sentry.Dsn != "" && strings.Contains(setting.Conf.Sentry.Env, os.Getenv("ENV")) { // 启用sentry时
		logger.Infof("sentry enabled")
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              setting.Conf.Sentry.Dsn,
			Debug:            setting.Conf.App.Debug,
			AttachStacktrace: false,
			Environment:      os.Getenv("ENV"),
		}); err != nil {
			logger.Errorf("sentry initialization failed: %v", err)
		}
		// 启用sentry
		engine.Use(gin.Recovery(), sentryGin.New(sentryGin.Options{Repanic: true}))
	} else { // sentry没有启用时，使用gin框架自身的Recovery来处理panic
		engine.Use(gin.Recovery())
	}

	logger.Infof("init sentry success")
}

func initPrometheus() {
	p := NewPrometheus("gin")
	p.Use(engine)
}

func initLogger() {
	logger = log.Sugar()
	logger.Info("logger init success")
}
