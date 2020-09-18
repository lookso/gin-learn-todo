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
	"gin-learn-todo/pkg/trace"
	"os"
)

func main() {
	fmt.Println("current pid:", os.Getpid())

	log.Init()

	if err := trace.NewTrace(); err != nil {
		log.Sugar().Errorf("ZinKin.Init() error(%v)", err)
		panic(err)
	}
	if err := db.Init(); err != nil {
		log.Sugar().Errorf("db.Init() error(%v)", err)
		panic(err)
	}
	defer db.Close()

	if err := redis.Init(); err != nil {
		log.Sugar().Errorf("redis.Init() error(%v)", err)
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
