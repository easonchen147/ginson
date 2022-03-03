package cache

import (
	"ginson/platform/cache"
	"github.com/go-redis/redis/v8"
)

type BaseCache struct {}

var baseCache = &BaseCache{}

func (b *BaseCache) client() *redis.Client {
	return cache.Client()
}
