package user

import (
	"context"
	"ginson/pkg/repository/mysql"
	"gorm.io/gorm"
)

type Db struct {
	*mysql.BaseDb
}

func NewDb() *Db {
	return &Db{BaseDb: mysql.NewBaseDb()}
}

func (q *Db) CreateUser(ctx context.Context, user *User) error {
	return q.Db().Create(user).Error
}

func (q *Db) GetUserById(ctx context.Context, userId uint) (*User, error) {
	var user User
	err := q.Db().First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *Db) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*User, error) {
	var user User
	err := q.Db().Where("open_id = ? and source = ? ", openId, source).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *Db) UpdateUserById(ctx context.Context, user *UserInfo) error {
	return q.Db().Model(&User{
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
