package redis

import (
	"fmt"
	"gin-learn-todo/app/setting"
	"github.com/go-redis/redis"
	"time"
)

var (
	Client *redis.Client
)

func init() {

	redisConf := setting.Init().Redis
	if redisConf != nil {
		Client = initRedisClientByConf(redisConf)
	}
}

// 初始化redis连接池
func initRedisClientByConf(redisConf *setting.Redis) (client *redis.Client) {
	options := &redis.Options{
		Addr:         redisConf.Host + ":" + redisConf.Port,
		DialTimeout:  time.Duration(redisConf.DialTimeoutMillisecond) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConf.ReadTimeoutMillisecond) * time.Millisecond,
		WriteTimeout: time.Duration(redisConf.WriteTimeoutMillisecond) * time.Millisecond,
		PoolSize:     redisConf.PoolSize,
	}

	client = redis.NewClient(options)
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf(pong)
	}
	return client
}
