package mongo

import (
	"ginson/platform/database"
)

type BaseMongo struct{}

func NewBaseMongo() *BaseMongo {
	return &BaseMongo{}
}

func (b *BaseMongo) Mgo() *database.MongoInstance {
	return database.Mongo()
}
