package cachedrepository

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type CachedRepository struct {
	repository Repository
	redis      *redis.Client
}

type Repository interface {
	GetURL(ctx context.Context, URL string) (string, error)
	ShortenURL(ctx context.Context, URL string, shortkey string) error
}

func NewCachedRepository(redis *redis.Client, repository Repository) *CachedRepository {
	return &CachedRepository{
		repository: repository,
		redis:      redis,
	}
}

func InitRedis(redisPort int, redisHost string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
}

func (r *CachedRepository) GetURL(ctx context.Context, URL string) (string, error) {
	shortenedURL, err := r.redis.Get(ctx, URL).Result()
	if errors.Is(err, redis.Nil) {

		shortenedURL, err = r.repository.GetURL(ctx, URL)
		if err != nil {
			err = fmt.Errorf("r.repository.GetURL(...): %w", err)
			return "", err
		}

		_, err = r.redis.Set(ctx, URL, shortenedURL, time.Hour).Result()
		if err != nil {
			err = fmt.Errorf("r.redis.Set(...): %w", err)
			return "", err
		}

		return shortenedURL, nil
	}
	if err != nil {
		err = fmt.Errorf("r.redis.Get(...): %w", err)
		return "", err
	}
	return shortenedURL, nil
}

func (r *CachedRepository) ShortenURL(ctx context.Context, URL string, shortkey string) error {
	err := r.repository.ShortenURL(ctx, URL, shortkey)
	if err != nil {
		err = fmt.Errorf("r.repository.ShortenURL(...): %w", err)
		return err
	}

	_, err = r.redis.Set(ctx, URL, shortkey, time.Hour).Result()
	if err != nil {
		err = fmt.Errorf("r.redis.Set(...): %w", err)
		return err
	}

	return nil
}
