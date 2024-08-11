package db

import (
	"context"
	"time"
)

var (
	DefaultCacheExpireTime = 1
)

func GetDataFromKey(key string) (interface{}, error) {
	ctx := context.Background()

	redisClient := DB.Redis

	value, err := redisClient.Get(ctx, key).Result()

	return value, err
}

func SaveDataToCache(key string, data interface{}) error {
	ctx := context.Background()

	redisClient := DB.Redis
	err := redisClient.Set(ctx, key, data, time.Duration(DefaultCacheExpireTime)*time.Hour).Err()
	return err
}

func SaveActiveTokenToCache(username string, data interface{}) error {
	ctx := context.Background()

	redisClient := DB.Redis
	key := "active_" + username
	err := redisClient.Set(ctx, key, data, time.Duration(DefaultCacheExpireTime)*time.Hour).Err()
	return err
}
