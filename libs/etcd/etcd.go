package etcd

import (
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli *clientv3.Client

var (
	LeaseKeepAliveFail = errors.New("lease keepAlive fail")
)

func Init() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	// 建立客户端成功
	fmt.Println("connect success")
}
func mustInit() error {
	if cli == nil {
		return errors.New("config is not init")
	}
	return nil
}