package shortener

import "url-shortener/domain/shortener/model"

type UsecaseInterface interface {
	Shorten(payload model.Payload) (*model.ShortURL, error)
	Redirect(shortURL string) (*string, error)
}
