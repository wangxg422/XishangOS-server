package initial

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.AppConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		logger.Info("redis connect ping response:", zap.String("pong", pong))
		global.RedisClient = client
	}
}
