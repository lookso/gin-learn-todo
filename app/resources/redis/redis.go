package redis

import (
	"gin-learn-todo/app/setting"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	client *redis.Client
)

func Init() {

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
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("init redis err %s", err)
	}
	log.Println("init redis success")
}

// Close ...
func Close() {
	client.Close()
}

func Client() *redis.Client {
	return client
}