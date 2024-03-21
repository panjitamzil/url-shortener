package route

import (
	"url-shortener/domain/shortener/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	shortenerHandler handler.Handler
}

func NewRoute(shortenerHandler handler.Handler) *Route {
	return &Route{
		shortenerHandler: shortenerHandler,
	}
}

func Routes(h *Route) *gin.Engine {
	r := gin.Default()
	r.POST("/shorten", h.shortenerHandler.Shorten)
	r.GET("/resolve/:url", h.shortenerHandler.Resolve)
	r.GET("/", h.shortenerHandler.Home)

	return r
}
