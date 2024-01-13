package main

import (
	config "url-shortener/cmd"
	"url-shortener/domain/shortener/handler"
	"url-shortener/domain/shortener/repository"
	"url-shortener/domain/shortener/usecase"
	"url-shortener/route"
)

func main() {
	// Load ENV
	err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	redis, err := config.InitRedis()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(redis)
	service := usecase.NewService(repo)
	h := handler.NewHandler(service)
	r := route.NewRoute(*h)

	apps := route.Routes(r)

	apps.Run(config.Cfg.App.Host + config.Cfg.App.Port)
}
