package user

import (
	"context"
	"ginson/pkg/log"
	"ginson/pkg/repository/mongo"
)

type Mongo struct {
	*mongo.BaseMongo
}

func NewMongo() *Mongo {
	return &Mongo{BaseMongo: mongo.NewBaseMongo()}
}

func (m *Mongo) AddUserCache(ctx context.Context, user *User) error {
	result, err := m.Mgo().Db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Info(ctx, "add user result: %v", result.InsertedID)
	return nil
}
