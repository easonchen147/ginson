package database

import (
	"errors"
	"fmt"
	"ginson/pkg/conf"
	"ginson/pkg/log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
func InitDB(cfg *conf.AppConfig) error {
	conns = make(map[string]*gorm.DB)
	for dbKey, dbConfig := range cfg.DbsConfig {
		conn, err := openConn(dbConfig.Uri, dbConfig.MaxIdleConn, dbConfig.MaxOpenConn)
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

func openConn(url string, idle, open int) (*gorm.DB, error) {
	newLogger := logger.New(Writer{}, logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true})
	openDB, err := gorm.Open(mysql.New(mysql.Config{DSN: url}), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	db, err := openDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(open)
	return openDB, nil
}

// Writer 记录SQL日志
type Writer struct{}

func (w Writer) Printf(format string, args ...interface{}) {
	log.Debug(fmt.Sprintf(format, args...))
}
