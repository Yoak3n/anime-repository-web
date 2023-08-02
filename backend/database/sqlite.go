package database

import (
	"fmt"
	"github/Yoak3n/anime-repository-web/package/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSqlite() *gorm.DB {
	dsn := fmt.Sprintf("data/db/sqlite.db")
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ERROR.Println(fmt.Sprintf("database connected err:%v", err))
	}
	if err != nil {
		logger.ERROR.Println(fmt.Sprintf("database connected err:%v", err))
	}
	return db
}
