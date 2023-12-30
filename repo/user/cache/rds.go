package cache

import (
	"context"
	"fmt"
	"ginson/api"
	"time"

	"github.com/easonchen147/foundation/cache"

	"ginson/pkg/util"
	"github.com/redis/go-redis/v9"
)

type Rds struct {
	client *redis.Client
}

func NewRds() *Rds {
	return &Rds{client: cache.Redis()}
}

func (c *Rds) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *Rds) SetUser(ctx context.Context, user *api.UserVO) error {
	return util.SetJsonCache(ctx, c.client, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *Rds) GetUser(ctx context.Context, userId uint) (*api.UserVO, error) {
	return util.GetByJsonCache[api.UserVO](ctx, c.client, c.getUserIdKey(userId))
}
