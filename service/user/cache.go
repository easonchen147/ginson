package user

import (
	"context"
)

type Cache interface {
	SetUser(ctx context.Context, user *UserVO) error
	GetUser(ctx context.Context, userId uint) (*UserVO, error)
}
