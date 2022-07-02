package models

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient  *redis.Client
	redisAddr    = "192.168.2.230:6379"
	redisTimeOut = 10 * time.Second
	redisPasswd  = "123456"
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    redisPasswd,
		DialTimeout: redisTimeOut,
	})
}

func RedisSetKey(key, value string) error {
	return RedisSetKeyWithTTL(key, value, 0)
}

func RedisSetKeyWithTTL(key, value string, ttl int64) error {
	return redisClient.Set(key, value, time.Duration(ttl)*time.Second).Err()
}

func RedisGetKey(key string) (string, error) {
	return redisClient.Get(key).Result()
}

func RedisDelKey(key string) error {
	return redisClient.Del(key).Err()
}
