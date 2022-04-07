package user

import (
	"context"
	"fmt"
	"ginson/pkg/repository/cache"
	"time"
)

type Cache struct {
	*cache.BaseCache
}

func NewCache() *Cache {
	return &Cache{BaseCache: cache.NewBaseCache()}
}

func (c *Cache) getUserIdKey(userId uint) string {
	return fmt.Sprintf("userId:%d", userId)
}

func (c *Cache) SetUser(ctx context.Context, user *UserInfo) error {
	return setJson(ctx, c.getUserIdKey(user.UserId), user, time.Hour)
}

func (c *Cache) GetUser(ctx context.Context, userId uint) (*UserInfo, error) {
	return getByJson[model.UserInfo](ctx, c.getUserIdKey(userId))
}
