package util

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetJsonCache(ctx context.Context, client *redis.Client, key string, obj any, ttl time.Duration) error {
	if client == nil {
		return errors.New("foundation client is nil")
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, string(jsonBytes), ttl).Err()
}

func GetByJsonCache[T any](ctx context.Context, client *redis.Client, key string) (*T, error) {
	if client == nil {
		return nil, errors.New("foundation client is nil")
	}
	jsonStr, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	obj := new(T)
	err = json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
