package usecase

import (
	"errors"
	"fmt"
	config "url-shortener/cmd"
	"url-shortener/domain/shortener"
	"url-shortener/domain/shortener/model"
	"url-shortener/lib"
)

type Service struct {
	repository shortener.RepositoryInterface
}

func NewService(repository shortener.RepositoryInterface) shortener.UsecaseInterface {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Shorten(payload model.Payload) (*model.ShortURL, error) {
	total, err := s.repository.GetKeys(payload.URL)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if total > 0 {
		return nil, errors.New("THE URL HAS BEEN SHORTENED")
	}

	shortURL := lib.GenerateShortKey()
	valid := lib.ValidateURL(payload.URL)
	if !valid {
		return nil, errors.New("INVALID URL")
	}

	key := fmt.Sprintf("shortener:%s", shortURL)

	err = s.repository.SetKey(key, payload.URL)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	updatedKey := fmt.Sprintf("shortener:%s:%s", shortURL, payload.URL)
	err = s.repository.SetKey(updatedKey, "1")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var data = model.ShortURL{
		Original: payload.URL,
		Short:    fmt.Sprintf("%s%s/resolve/%s", config.Cfg.App.Host, config.Cfg.App.Port, shortURL),
	}

	return &data, nil
}

func (s *Service) Redirect(shortURL string) (*string, error) {
	key := fmt.Sprintf("shortener:%s", shortURL)

	result, err := s.repository.GetKeys(key)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if result == 0 {
		return nil, errors.New("URL IS NOT EXIST")
	}

	url, err := s.repository.GetKey(key)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &url, nil
}
