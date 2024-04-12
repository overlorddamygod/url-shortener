package service

import (
	"errors"

	"github.com/overlorddamygod/url-shortener/utils"
)

type ShortenerService struct {
	urlMap map[string]string
}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{
		urlMap: make(map[string]string),
	}
}

func (s *ShortenerService) GetShortCode(longURL string) (string, error) {
	if !utils.IsValidURL(longURL) {
		return "", errors.New("invalid url")
	}

	uniqueIdentifier := utils.GetHash(longURL)
	shortCode := utils.EncodeBase62(uniqueIdentifier, 8)

	s.urlMap[shortCode] = longURL

	return shortCode, nil
}

func (s *ShortenerService) GetLongUrl(shortCode string) (string, error) {
	longUrl, ok := s.urlMap[shortCode]
	if !ok {
		return "", errors.New("short code not found")
	}

	return longUrl, nil
}
