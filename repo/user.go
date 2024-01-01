package repo

import (
	"context"
	"ginson/models/api"
	"ginson/models/entity"
)

type UserCache interface {
	SetUser(ctx context.Context, user *api.UserVO) error
	GetUser(ctx context.Context, userId uint) (*api.UserVO, error)
}

type UserDb interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserById(ctx context.Context, userId uint) (*entity.User, error)
	FindByOpenIdAndSource(ctx context.Context, openId, source string) (*entity.User, error)
	UpdateUserById(ctx context.Context, user *entity.User) error
}
