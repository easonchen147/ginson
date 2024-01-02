package userrepo

import (
	"context"
	"ginson/api"
	"ginson/model"
)

type Cache interface {
	SetUser(ctx context.Context, user *api.UserVO) error
	GetUser(ctx context.Context, userId uint) (*api.UserVO, error)
}

type Db interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, userId uint) (*model.User, error)
	FindByOpenIdAndSource(ctx context.Context, openId, source string) (*model.User, error)
	UpdateUserById(ctx context.Context, user *model.User) error
}
