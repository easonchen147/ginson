package database

import (
	"errors"
	"fmt"
	"ginson/foundation/cfg"
	"ginson/foundation/log"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	conns map[string]*gorm.DB
)

func DB(dbName ...string) *gorm.DB {
	if db == nil && len(conns) == 0 {
		panic(errors.New("mysql is not ready"))
	}
	if len(dbName) > 0 {
		if conn, ok := conns[dbName[0]]; ok {
			return conn
		}
	}
	return db
}

// InitDB 初始化数据库
func InitDB(cfg *cfg.AppConfig) error {
	conns = make(map[string]*gorm.DB)
	for dbKey, dbConfig := range cfg.DbsConfig {
		conn, err := openConn(dbConfig.Uri, dbConfig.MaxIdleConn, dbConfig.MaxOpenConn, dbConfig.ConnectLifeTime, dbConfig.ConnectIdleTime)
		if err != nil {
			return fmt.Errorf("open connection failed, error: %s", err.Error())
		}
		conns[dbKey] = conn
		if dbKey == "default" {
			db = conn
		}
	}
	return nil
}

func openConn(url string, idle, open, lifeTime, idleTime int) (*gorm.DB, error) {
	newLogger := zapgorm2.New(log.SqlLogger)
	newLogger.SetAsDefault()
	openDB, err := gorm.Open(mysql.New(mysql.Config{DSN: url}), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, err
	}

	db, err := openDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(open)
	if lifeTime > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(lifeTime))
	}
	if idleTime > 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(idleTime))
	}
	return openDB, nil
}
