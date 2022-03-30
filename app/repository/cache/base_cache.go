package cache

import (
	"context"
	"encoding/json"
	"ginson/platform/cache"
	"github.com/go-redis/redis/v8"
	"time"
)

type BaseCache struct{}

var baseCache = &BaseCache{}

func (b *BaseCache) redis() *redis.Client {
	return cache.Redis()
}

func setJson(ctx context.Context, key string, obj any, ttl time.Duration) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return cache.Redis().Set(ctx, key, string(jsonBytes), ttl).Err()
}

func getByJson[T any](ctx context.Context, key string) (*T, error) {
	jsonStr, err := cache.Redis().Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	obj := new(T)
	err = json.Unmarshal([]byte(jsonStr), &obj)
	return obj, err
}
