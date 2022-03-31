package mysql

import (
	"context"
	"ginson/app/model"
)

type UserQuery struct {
	*BaseQuery
}

var userQuery = &UserQuery{BaseQuery: baseQuery}

func GetUserQuery() *UserQuery {
	return userQuery
}

func (q *UserQuery) CreateUser(ctx context.Context, user *model.User) error {
	return q.db().Create(user).Error
}

func (q *UserQuery) GetUserById(ctx context.Context, userId uint) (*model.User, error) {
	var user model.User
	err := q.db().First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *UserQuery) FindByOpenIdAndSource(ctx context.Context, openId, source string) (*model.User, error) {
	var user model.User
	err := q.db().Where("openId = ? and source = ? ", openId, source).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
