package user

import (
	"context"
	"github.com/easonchen147/foundation/database"

	"gorm.io/gorm"
)

type RepositoryDb struct {
	db *gorm.DB
}

func NewRepositoryDb() *RepositoryDb {
	return &RepositoryDb{db: database.DB()}
}

func (r *RepositoryDb) CreateUser(ctx context.Context, user *User) error {
	return r.db.Create(user).Error
}

func (r *RepositoryDb) GetUserById(ctx context.Context, userId uint) (*User, error) {
	var result User
	err := r.db.First(&result, userId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *RepositoryDb) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*User, error) {
	var result User
	err := r.db.Where("open_id = ? and source = ? ", openId, source).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *RepositoryDb) UpdateUserById(ctx context.Context, user *User) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"age":      user.Age,
		"gender":   user.Gender,
	}).Error
}
