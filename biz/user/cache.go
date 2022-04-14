package user

import (
	"context"
	"fmt"
	"ginson/foundation/cache"
	"ginson/pkg/util"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewCache() *Cache {
	return &Cache{client: cache.Redis()}
}

func (c *Cache) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *Cache) SetUser(ctx context.Context, user *User) error {
	return util.SetJsonCache(ctx, c.client, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *Cache) GetUser(ctx context.Context, userId uint) (*User, error) {
	return util.GetByJsonCache[User](ctx, c.client, c.getUserIdKey(userId))
}
