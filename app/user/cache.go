package user

import (
	"context"
	"fmt"
	"ginson/pkg/util"
	"ginson/platform/database"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache struct {
	client *redis.Client
}

func NewCache() *Cache {
	return &Cache{client: database.Redis()}
}

func (c *Cache) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *Cache) SetUser(ctx context.Context, user *UserInfo) error {
	return util.SetJsonCache(ctx, c.client, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *Cache) GetUser(ctx context.Context, userId uint) (*UserInfo, error) {
	return util.GetByJsonCache[UserInfo](ctx, c.client, c.getUserIdKey(userId))
}
