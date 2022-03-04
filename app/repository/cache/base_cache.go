package cache

import (
	"ginson/platform/cache"
	"github.com/go-redis/redis/v8"
)

type BaseCache struct{}

var baseCache = &BaseCache{}

func (b *BaseCache) redis() *redis.Client {
	return cache.Redis()
}
