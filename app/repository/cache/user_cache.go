package cache

import (
	"context"
	"fmt"
	"ginson/app/models"
	"time"
)

type UserCache struct {
	*BaseCache
}

var userCache = &UserCache{BaseCache: baseCache}

func GetUserCache() *UserCache {
	return userCache
}

func (c *UserCache) AddUserCache(ctx context.Context, user *models.User) error {
	return c.client().Set(ctx, fmt.Sprintf("userId:%d", user.Id), user, time.Hour).Err()
}
