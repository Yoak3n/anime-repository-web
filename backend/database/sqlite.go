package database

import (
	"fmt"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSqlite() *gorm.DB {
	util.CreateDirNotExists("data/db")
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
