package cache

import (
	"context"
	"ginson/api"
)

type RepoCache interface {
	SetUser(ctx context.Context, user *api.UserVO) error
	GetUser(ctx context.Context, userId uint) (*api.UserVO, error)
}
