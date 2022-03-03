package mysql

import (
	"context"
	"errors"
	"ginson/app/models"
	"gorm.io/gorm"
)

type UserQuery struct{
	*BaseQuery
}

var userQuery = &UserQuery{BaseQuery:baseQuery}

func GetUserQuery() *UserQuery {
	return userQuery
}

func (q *UserQuery) CreateUser(ctx context.Context, user *models.User) error {
	return q.db().Create(user).Error
}

func (q *UserQuery) IsUserEmailExisted(ctx context.Context, email string) (bool, error) {
	var user models.User
	err := q.db().Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (q *UserQuery) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := q.db().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}