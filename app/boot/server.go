/*
@Time : 2019-11-24 11:00 
@Author : Tenlu
@File : server
@Software: GoLand
*/
package boot

import (
	"context"
	db "gin-learn-todo/app/resources/mysql"
	"gin-learn-todo/app/resources/redis"
	"gin-learn-todo/app/routers"
	"gin-learn-todo/app/setting"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func New() {

	if err := db.Init(); err != nil {
		log.Fatalf("mysql.Init() error(%v)", err)
		return
	}
	defer db.Close()
	redis.Init()
	defer redis.Close()

	InitSentry()

	InitGin()
}

func InitGin() {
	router := routers.SetUpRouter()

	// 初始化http server
	s := &http.Server{
		Addr:              setting.Conf.ApiServer.ListenAddr,
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
		//s.logger.Debug("server listen at ", s.config.Listen)
		// service connections
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("http.server start success,listen.port%s", setting.Conf.ApiServer.ListenAddr)

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
		log.Fatal("server Shutdown error: ", err)
	}
	log.Println("server exiting success")
}

func InitSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              setting.Conf.Sentry.Dsn,
		AttachStacktrace: true,
	}); err != nil {
		log.Fatalf("sentry initialization failed: %v", err)
	}
	log.Println("init sentry success")
}

func AddRouters(f func(r *gin.Engine)) {
	var engine *gin.Engine
	f(engine)
}
