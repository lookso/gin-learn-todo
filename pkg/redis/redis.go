package redis

import (
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/setting"
	"github.com/go-redis/redis"
	"time"
)

var (
	client *redis.Client
)

const (
	ErrNil = redis.Nil
)

func Init() (err error) {
	redisConf := setting.Conf.Redis
	if redisConf == nil {
		panic("init redis error")
	}
	options := &redis.Options{
		Addr:         redisConf.Addr,
		DialTimeout:  time.Duration(redisConf.DialTimeoutMillisecond) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConf.ReadTimeoutMillisecond) * time.Millisecond,
		WriteTimeout: time.Duration(redisConf.WriteTimeoutMillisecond) * time.Millisecond,
		PoolSize:     redisConf.PoolSize,
	}

	client = redis.NewClient(options)
	_, err = client.Ping().Result()
	if err != nil {
		log.Sugar().Errorf("init redis err %s", err)
		return err
	}
	log.Sugar().Info("init redis success")
	return nil
}

// Close ...
func Close() {
	client.Close()
}

func Client() *redis.Client {
	return client
}
