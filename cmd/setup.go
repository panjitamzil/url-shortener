package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitRedis() (*redis.Client, error) {
	add := fmt.Sprintf("%s%s", Cfg.Redis.Address, Cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: add,
		DB:   0,
	})

	pong, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(pong, "Redis connected!")
	return rdb, nil
}
