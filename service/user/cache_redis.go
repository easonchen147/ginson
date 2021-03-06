package user

import (
	"context"
	"fmt"
	"ginson/pkg/util"
	"github.com/easonchen147/foundation/cache"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheRedis struct {
	client *redis.Client
}

func NewCacheRedis() *CacheRedis {
	return &CacheRedis{client: cache.Redis()}
}

func (c *CacheRedis) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *CacheRedis) SetUser(ctx context.Context, user *UserVO) error {
	return util.SetJsonCache(ctx, c.client, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *CacheRedis) GetUser(ctx context.Context, userId uint) (*UserVO, error) {
	return util.GetByJsonCache[UserVO](ctx, c.client, c.getUserIdKey(userId))
}
