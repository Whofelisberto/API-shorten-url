package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

type RedisStore struct {
	rdb *redis.Client
}

func NewStore(rdb *redis.Client) Store {
	return &RedisStore{rdb: rdb}
}

func (s *RedisStore) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string

	for i := 0; i < 5; i++ {
		code = genCode()
		err := s.rdb.HGet(ctx, "shorten_url", code).Err()
		if errors.Is(err, redis.Nil) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to get url from shorten_url hashmap: %w", err)
		}
	}

	if err := s.rdb.HSet(ctx, "shorten_url", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in shorten_url hashmap: %w", err)
	}
	return code, nil
}

func (s *RedisStore) GetFullURL(ctx context.Context, code string) (string, error) {
	fullURL, err := s.rdb.HGet(ctx, "shorten_url", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get url from shorten_url hashmap: %w", err)
	}
	return fullURL, nil
}