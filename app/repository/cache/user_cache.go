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

func (c *UserCache) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *UserCache) SetUser(ctx context.Context, user *model.UserInfo) error {
	return setJson(ctx, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *UserCache) GetUser(ctx context.Context, userId uint) (*model.UserInfo, error) {
	return getByJson[model.UserInfo](ctx, c.getUserIdKey(userId))
}
