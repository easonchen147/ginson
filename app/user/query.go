package user

import (
	"context"
	"ginson/platform/database"
	"gorm.io/gorm"
)

type Query struct {
	db *gorm.DB
}

func NewQuery() *Query {
	return &Query{db: database.DB()}
}

func (q *Query) CreateUser(ctx context.Context, user *User) error {
	return q.db.Create(user).Error
}

func (q *Query) GetUserById(ctx context.Context, userId uint) (*User, error) {
	var user User
	err := q.db.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *Query) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*User, error) {
	var user User
	err := q.db.Where("open_id = ? and source = ? ", openId, source).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *Query) UpdateUserById(ctx context.Context, user *UserInfo) error {
	return q.db.Model(&User{
		Model: gorm.Model{
			ID: user.UserId,
		},
	}).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"age":      user.Age,
		"gender":   user.Gender,
	}).Error
}
