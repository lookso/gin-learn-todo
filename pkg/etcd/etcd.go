package etcd

import (
	"errors"
	"gin-learn-todo/pkg/log"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var Cli *clientv3.Client

var (
	LeaseKeepAliveFail = errors.New("lease keepAlive fail")
)

func Init() error {
	var err error
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:12379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Sugar().Fatalf("init etcd err %s", err)
		return err
	}
	// 建立客户端成功
	log.Sugar().Info("init etcd success")
	return nil
}
func MustInit() error {
	if Cli == nil {
		return errors.New("config is not init")
	}
	return nil
}
