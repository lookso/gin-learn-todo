package etcd

import (
	"errors"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var cli *clientv3.Client

var (
	LeaseKeepAliveFail = errors.New("lease keepAlive fail")
)

func Init() error{
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("init etcd err %s", err)
		return err
	}
	// 建立客户端成功
	log.Printf("init etcd success")
	return nil
}
func mustInit() error {
	if cli == nil {
		return errors.New("config is not init")
	}
	return nil
}
