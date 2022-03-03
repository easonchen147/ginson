package mysql

import (
	"ginson/platform/database"
	"gorm.io/gorm"
)

type BaseQuery struct {}

var baseQuery = &BaseQuery{}

func (b *BaseQuery) db(dbName ...string) *gorm.DB {
	return database.DB(dbName...)
}