package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/wangxg422/XishangOS-backend/config"
)

var (
	AppConfig   config.AppConfig
	RedisClient *redis.Client
)
