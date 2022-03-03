package mongo

import (
	"context"
	"fmt"
	"ginson/app/models"
	"ginson/pkg/log"
)

type UserMongo struct {
	*BaseMongo
}

var userMongo = &UserMongo{BaseMongo: baseMongo}

func GetUserMongo() *UserMongo {
	return userMongo
}

func (m *UserMongo) AddUserCache(ctx context.Context, user *models.User) error {
	result, err := m.mg().Db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("add user result: %v", result.InsertedID))
	return nil
}