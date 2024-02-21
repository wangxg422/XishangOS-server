package jwt

import (
	"context"
	"github.com/wangxg422/XishangOS-backend/global"
)

func GetRedisJWT(username string) (string, error) {
	return global.RedisClient.Get(context.Background(), global.AppConfig.Jwt.RedisKey+username).Result()
}

func SetRedisJWT(token, username string) error {
	dr, err := ParseDuration(global.AppConfig.Jwt.ExpiresTime)
	if err != nil {
		return err
	}

	return global.RedisClient.Set(context.Background(), global.AppConfig.Jwt.RedisKey+username, token, dr).Err()
}

func InBlackList(token string) (bool, error) {
	result, err := global.RedisClient.HGet(context.Background(), global.AppConfig.Jwt.BlackListKey, token).Result()
	return result != "", err
}

func ToBlackList(token string) (bool, error) {
	result, err := global.RedisClient.HSet(context.Background(), global.AppConfig.Jwt.BlackListKey, token).Result()
	return result != 0, err
}
