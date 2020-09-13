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
		log.Sugar().Errorf("mysql.Init() error(%v)", err)
		return
	}
	defer db.Close()

	if err := redis.Init(); err != nil {
		log.Sugar().Errorf("redis.Init() error(%v)", err)
		return
	}
	defer redis.Close()

	if err := etcd.Init(); err != nil {
		log.Sugar().Errorf("etcd.Init() error(%v)", err)
		return
	}

	boot.NewServer()
}
