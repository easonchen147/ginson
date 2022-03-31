package mongo

import (
	"ginson/platform/database"
)

type BaseMongo struct{}

var baseMongo = &BaseMongo{}

func (b *BaseMongo) mg() *database.MongoInstance {
	return database.Mongo()
}
