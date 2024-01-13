package repository

import (
	"url-shortener/domain/shortener"

	"github.com/go-redis/redis/v8"
)

type Repo struct {
	redis *redis.Client
}

func NewRepo(redis *redis.Client) shortener.RepositoryInterface {
	return &Repo{
		redis: redis,
	}
}
