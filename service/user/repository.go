package user

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, userId uint) (*User, error)
	FindByOpenIdAndSource(ctx context.Context, openId, source string) (*User, error)
	UpdateUserById(ctx context.Context, user *User) error
}
