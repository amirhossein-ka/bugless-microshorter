package service

import "errors"

// TODO: add validation for urls ? idk

// AddUrl implements Service.
func (s *srvImpl) AddUrl(url string) (string, error) {
	result, err := s.newUrlRpc([]string{url})
	if err != nil {
		return "", err
	}
	if len(result) <= 0 {
		return "", errors.New("no results")
	}

	return result[0], nil
}

// GetFullURL implements Service.
func (s *srvImpl) GetFullURL(key string) (string, error) {
	// check if key exists in cache
	if v, ok := s.cache.Get(key); ok {
		return v, nil
	}

	result, err := s.getUrlRpc([]string{key})
	if err != nil {
		return "", err
	}
	if len(result) <= 0 {
		return "", errors.New("no results")
	}

	return result[0], nil
}
