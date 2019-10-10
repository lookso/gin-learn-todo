package main

import (
	"context"
	"fmt"
	"gin-sourcecode-learn/router"
	"gin-sourcecode-learn/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

// 参考:https://www.jianshu.com/p/35addb4de300
// https://studygolang.com/articles/23097

func main() {
	fmt.Println("Gin version", gin.Version)

	goVersion, _ := utils.GetMinVer(runtime.Version())
	if goVersion > 8 {
		fmt.Println(runtime.Version()) // 获取go 当前版本号 // go1.12.9
	}
	utils.DebugPrint(`the chinese people is nb`)

	engine := router.InitRouter()
	//engine.Run(":8989")

	srv := &http.Server{
		Addr:    ":8989",
		Handler: engine,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
