package db

import (
	"context"
	"ginson/model"
	"github.com/easonchen147/foundation/db"

	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func NewDb() *Db {
	return &Db{db: db.DB()}
}

func (r *Db) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.Create(user).Error
}

func (r *Db) GetUserById(ctx context.Context, userId uint) (*model.User, error) {
	var result model.User
	err := r.db.First(&result, userId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Db) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*model.User, error) {
	var result model.User
	err := r.db.Where("open_id = ? and source = ? ", openId, source).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Db) UpdateUserById(ctx context.Context, user *model.User) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"age":      user.Age,
		"gender":   user.Gender,
	}).Error
}
