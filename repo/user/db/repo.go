package db

import (
	"context"
	"ginson/model"
)

type Repodb interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, userId uint) (*model.User, error)
	FindByOpenIdAndSource(ctx context.Context, openId, source string) (*model.User, error)
	UpdateUserById(ctx context.Context, user *model.User) error
}
