package cache

import (
	"context"
	"fmt"
	"ginson/app/model"
	"time"
)

type UserCache struct {
	*BaseCache
}

var userCache = &UserCache{BaseCache: baseCache}

func GetUserCache() *UserCache {
	return userCache
}

func (c *UserCache) AddUserCache(ctx context.Context, user *model.User) error {
	return c.redis().Set(ctx, fmt.Sprintf("userId:%d", user.Id), user, time.Hour).Err()
}
