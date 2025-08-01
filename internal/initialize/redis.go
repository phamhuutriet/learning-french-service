package initialize

import (
	"context"
	"fmt"
	"learning-french-service/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("failed to connect to redis", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("redis connected", zap.String("host", fmt.Sprintf("%s:%d", r.Host, r.Port)))
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	rdb := global.Rdb
	rdb.Set(ctx, "score", 100, 0)
	val, err := rdb.Get(ctx, "score").Result()
	if err != nil {
		global.Logger.Error("failed to get key", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("redis get", zap.String("score", val))
}
