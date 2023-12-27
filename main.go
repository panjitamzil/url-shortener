package main

import (
	config "url-shortener/cmd"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load ENV
	err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	_, err = config.InitRedis()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// repo := shortener.NewRedisRepository(rdb)
	// service := shortener.NewShortURLService(repo)

	// usecase := shortener.NewShortURLUsecase(repo, service)

	// shortener.NewHandler(router, usecase)

	router.Run("localhost:8080")
}
