package cache

import (
	"context"
	"fmt"
	"github.com/gookit/color"
	"github.com/redis/go-redis/v9"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"time"
)

var ctx = context.Background()

var redisClient *redis.Client

func InitRedis(config *config.Config) (*redis.Client, error) {
	uri := fmt.Sprintf("redis://:%s@%s:%d", config.Redis.Password, config.Redis.Host, config.Redis.Port)
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	opt.PoolSize = config.Redis.RedisPoolSize
	opt.DB = config.Redis.DB
	client := redis.NewClient(opt)
	rctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err = client.Ping(rctx).Result()
	if err != nil {
		color.Redln("初始化Redis失败", err)
		return nil, err
	}
	redisClient = client
	return client, nil
}

func GetRedis() *redis.Client {
	if redisClient == nil {
		initRedis, err := InitRedis(config.NewConfig())
		if err != nil {
			return nil
		}
		return initRedis
	}
	return redisClient
}
