package mysql

import (
	"ginson/platform/database"
	"gorm.io/gorm"
)

type BaseDb struct{}

func NewBaseDb() *BaseDb {
	return &BaseDb{}
}

func (b *BaseDb) Db(dbName ...string) *gorm.DB {
	return database.DB(dbName...)
}
