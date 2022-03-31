package mysql

import (
	"ginson/platform/database"
	"gorm.io/gorm"
)

type BaseDb struct{}

var baseDb = &BaseDb{}

func (b *BaseDb) db(dbName ...string) *gorm.DB {
	return database.DB(dbName...)
}
