/*
@Time : 2019-11-17 14:40 
@Author : Tenlu
@File : service
@Software: GoLand
*/
package boot

import (
	"context"
	"fmt"
	"gin-learn-todo/app/conf"
	"gin-learn-todo/app/router"
	"gin-learn-todo/app/rpc"
	"gin-learn-todo/app/utils"
	"gin-learn-todo/app/utils/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"
)

type Server struct {
	config conf.Config

	logger *zap.SugaredLogger
	gin    *gin.Engine
	http   *http.Server
	rpc    *rpc.Service
}

func New() *Server {
	var cnf conf.Config
	srv := &Server{
		config: cnf,
	}
	srv.InitGin()
	srv.InitLogger()
	srv.InitRpc()

	return srv
}

func (s *Server) InitRpc() {
	rpcSrv := rpc.Service{}
	rpcSrv.Port = 90
	rpcSrv.StartRpc()
}

func (s *Server) InitLogger() {

	var err error
	// 初始化Logger
	if s.logger, err = log.New(&s.config.Log); err != nil {
		fmt.Println("config file unmarshal err:", err)
		os.Exit(1)
	}

	s.logger.Debug("Logger Init Success")
}

func (s *Server) InitGin() {
	const abortIndex int8 = math.MaxInt8 / 2
	fmt.Println("最多Handles个数:", abortIndex) // 63

	fmt.Println("Gin version", gin.Version)

	goVersion, _ := utils.GetMinVer(runtime.Version())
	if goVersion > 8 {
		fmt.Println(runtime.Version()) // 获取go 当前版本号 // go1.12.9
	}
	utils.DebugPrint(`the app run port:8989,` + ` the current pid is:` + strconv.Itoa(os.Getpid()))

	s.gin = router.InitRouter()

	//engine.Run(":8989")

	s.http = &http.Server{
		Addr:              ":8989",
		Handler:           s.gin,
		ReadTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 60,
	}

	go func() {
		// 服务连接
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.logger.Debug("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.http.Shutdown(ctx); err != nil {
		s.logger.Fatal("Server Shutdown:", err)
	}
	s.logger.Debug("Server exiting")
}
