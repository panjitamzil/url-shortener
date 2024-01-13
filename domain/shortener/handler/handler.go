package handler

import (
	"errors"
	"net/http"
	"url-shortener/domain/shortener"
	"url-shortener/domain/shortener/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase shortener.UsecaseInterface
}

func NewHandler(usecase shortener.UsecaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Shorten(c *gin.Context) {
	var payload model.Payload

	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := h.usecase.Shorten(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *Handler) Resolve(c *gin.Context) {
	shortURL := c.Param("url")

	url, err := h.usecase.Redirect(shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if url == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("URL IS NIL"),
		})
		return
	}

	c.Redirect(http.StatusFound, *url)
}
