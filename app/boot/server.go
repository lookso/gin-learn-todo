/*
@Time : 2019-11-24 11:00 
@Author : Tenlu
@File : server
@Software: GoLand
*/
package boot

import (
	"fmt"
	"gin-learn-todo/app/public/zaplog"
	"gin-learn-todo/app/setting"
	"os"
)

func Init() {
	if err := setting.Init(); err != nil {
		panic(err)
	}

	// 初始化Logger
	logger, err := zaplog.New(setting.Conf.LogConfig)
	if err != nil {
		fmt.Println("config file unmarshal err:", err)
		os.Exit(1)
	}
	logger.Info("Logger Init Success")

	//if err := mysql.Init(); err != nil {
	//	log.Error("dao.Init() error(%v)", err)
	//	return
	//}
	//if err := service.Init(); err != nil {
	//	log.Error("service.Init() error(%v)", err)
	//	return
	//}
	//
	//if err!=redis.Init();err!=nil{
	//
	//}

}
