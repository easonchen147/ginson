package user

import (
	"context"
	"fmt"
	"ginson/pkg/util"
	"ginson/platform/database"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func newCache() *cache {
	return &cache{client: database.Redis()}
}

func (c *cache) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *cache) SetUser(ctx context.Context, user *Info) error {
	return util.SetJsonCache(ctx, c.client, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *cache) GetUser(ctx context.Context, userId uint) (*Info, error) {
	return util.GetByJsonCache[Info](ctx, c.client, c.getUserIdKey(userId))
}
