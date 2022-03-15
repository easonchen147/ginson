package mongo

import (
	"context"
	"ginson/app/model"
	"ginson/pkg/log"
)

type UserMongo struct {
	*BaseMongo
}

var userMongo = &UserMongo{BaseMongo: baseMongo}

func GetUserMongo() *UserMongo {
	return userMongo
}

func (m *UserMongo) AddUserCache(ctx context.Context, user *model.User) error {
	result, err := m.mg().Db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Info("add user result: %v", result.InsertedID)
	return nil
}
