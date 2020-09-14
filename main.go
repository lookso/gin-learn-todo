/*
@Time : 2019-08-31 22:19
@Author : Tenlu
@File : api
@Software: GoLand
*/
package main

import (
	"gin-learn-todo/boot"
	"gin-learn-todo/pkg/etcd"
	"gin-learn-todo/pkg/log"
	db "gin-learn-todo/pkg/mysql"
	"gin-learn-todo/pkg/redis"
)

func main() {
	log.Init()

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

	if err := etcd.Init(); err != nil {
		log.Sugar().Errorf("etcd.Init() error(%v)", err)
		panic(err)
	}
	//go grpc.Run()
	//go grpc.Gateway()
	boot.NewServer()
}
