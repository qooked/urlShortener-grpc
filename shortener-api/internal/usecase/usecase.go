package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"net/url"
)

type Usecase struct {
	cachedRepository cachedRepository
	baseURL          string
}

type cachedRepository interface {
	ShortenURL(ctx context.Context, URL string, shortkey string) error
	GetURL(ctx context.Context, URL string) (string, error)
}

func NewUsecase(baseURL string, cachedRepository cachedRepository) *Usecase {
	return &Usecase{
		cachedRepository: cachedRepository,
		baseURL:          baseURL,
	}
}

func (u *Usecase) ShortenURL(ctx context.Context, URL string) (string, error) {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		err = fmt.Errorf("url.ParseRequestURI(...): %w", err)
		return "", err
	}

	shortkey := generateShortKey(URL)
	shortenedURL, err := url.JoinPath(u.baseURL, shortkey)
	if err != nil {
		err = fmt.Errorf("url.JoinPath(...): %w", err)
		return "", err
	}

	err = u.cachedRepository.ShortenURL(ctx, URL, shortenedURL)
	if err != nil {
		err = fmt.Errorf("u.cachedRepository.ShortenURL(...): %w", err)
		return "", err
	}

	return shortenedURL, nil
}
func (u *Usecase) GetURL(ctx context.Context, URL string) (string, error) {
	return u.cachedRepository.GetURL(ctx, URL)
}

func generateShortKey(URL string) string {
	crcHash := crc32.ChecksumIEEE([]byte(URL))

	base64Hash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", crcHash)))

	return base64Hash
}
