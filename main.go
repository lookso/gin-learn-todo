/*
@Time : 2019-08-31 22:19
@Author : Tenlu
@File : api
@Software: GoLand
*/
package main

import (
	"gin-learn-todo/boot"
	"gin-learn-todo/libs/etcd"
	db "gin-learn-todo/libs/mysql"
	"gin-learn-todo/libs/redis"
	"log"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("mysql.Init() error(%v)", err)
		return
	}
	defer db.Close()

	if err := redis.Init(); err != nil {
		log.Fatalf("redis.Init() error(%v)", err)
		return
	}
	defer redis.Close()

	if err := etcd.Init(); err != nil {
		log.Fatalf("etcd.Init() error(%v)", err)
		return
	}
	boot.NewServer()
}
