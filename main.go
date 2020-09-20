/*
@Time : 2019-08-31 22:19
@Author : Tenlu
@File : api
@Software: GoLand
*/
package main

import (
	"fmt"
	"gin-learn-todo/boot"
	"gin-learn-todo/pkg/grpc"
	"gin-learn-todo/pkg/log"
	db "gin-learn-todo/pkg/mysql"
	"gin-learn-todo/pkg/redis"
	"os"
)

func main() {
	fmt.Println("current pid:", os.Getpid())

	log.Init()

	//if err := trace.NewTrace(); err != nil {
	//	log.Sugar().Errorf("ZinKin init error(%v)", err)
	//	panic(err)
	//}
	if err := db.Init(); err != nil {
		log.Sugar().Errorf("db init error(%v)", err)
		panic(err)
	}
	defer db.Close()

	if err := redis.Init(); err != nil {
		log.Sugar().Errorf("redis init error(%v)", err)
		panic(err)
	}
	defer redis.Close()

	//if err := etcd.Init(); err != nil {
	//	log.Sugar().Errorf("etcd.Init() error(%v)", err)
	//	panic(err)
	//}
	go grpc.Run()
	go grpc.Gateway()
	boot.NewServer()

}
