/*
@Time : 2019-11-17 14:40 
@Author : Tenlu
@File : service
@Software: GoLand
*/
package boot

import (
	"context"
	"gin-learn-todo/app/router"
	"gin-learn-todo/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var logger *zap.SugaredLogger

type Server struct {

}

//
//http   *http.Server
//rpc    *rpc.Service

//config conf.Config

func NewSrv() *Server {
	srv := &Server{}

	return srv

}

func (s *Server) Run() {
	s.InitGin()
}

func (s *Server) InitGin() {

	utils.DebugPrint(`the app run port:9001` + `, the current pid is:` + strconv.Itoa(os.Getpid()))

	rs := gin.New()

	//if s.config.Debug {
	gin.SetMode(gin.DebugMode)
	//} else {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	engine := router.InitRouter(rs)

	srv := &http.Server{
		Addr:              ":9001",
		Handler:           engine,
		ReadTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 60,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Debug("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Debug("Server exiting")
}
