package service

import (
	"math/rand"
	"time"
	"url-shortener/store"
)

type UrlService struct {
	store *store.MemoryStore
}

func NewUrlService(store *store.MemoryStore) *UrlService {
	return &UrlService{store}
}

func (u *UrlService) ShortenUrl(url string) string {
	code := generateCode()
	u.store.Save(code, url)
	return code
}

func (s *UrlService) Resolve(code string) (string, bool) {
	return s.store.Get(code)
}

func generateCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(code))]
	}
	return string(code)

}
