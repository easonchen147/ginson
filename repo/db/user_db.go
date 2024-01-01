package db

import (
	"context"
	"ginson/models/entity"
	"github.com/easonchen147/foundation/db"

	"gorm.io/gorm"
)

type UserDb struct {
	db *gorm.DB
}

func NewUserDb() *UserDb {
	return &UserDb{db: db.DB()}
}

func (r *UserDb) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserDb) GetUserById(ctx context.Context, userId uint) (*entity.User, error) {
	var result entity.User
	err := r.db.First(&result, userId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserDb) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*entity.User, error) {
	var result entity.User
	err := r.db.Where("open_id = ? and source = ? ", openId, source).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserDb) UpdateUserById(ctx context.Context, user *entity.User) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"age":      user.Age,
		"gender":   user.Gender,
	}).Error
}
