package service

import (
	"context"
	"ush/internal/shortener/repository"
)

// should i use interfaces... ?

type (
	ShortenerService interface {
		AddUrl(url string) (string, error)
		GetUrl(key string) (string, error)

		// Stop is only a simple wrapper to close connections to db since
		// we dont have access to repository in controller (:
		Stop() error
	}

	serviceImpl struct {
		repo repository.Repository
	}
)

// Stop implements ShortenerService.
func (s *serviceImpl) Stop() error {
	return s.repo.Stop(context.Background())
}

// tbh im to tired to make different files for these...

func (s *serviceImpl) GetUrl(key string) (string, error) {
	url, err := s.repo.Get(context.Background(), key)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *serviceImpl) AddUrl(url string) (string, error) {
	key := randomString(10)
	_, err := s.repo.Create(context.Background(), url, key)
	if err != nil {
		// TODO: check if error is duplicate thing to recreate key
		return "", err
	}

	return key, nil
}

func New(repo repository.Repository) ShortenerService {
	return &serviceImpl{repo: repo}
}
