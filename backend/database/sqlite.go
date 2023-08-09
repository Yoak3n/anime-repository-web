package database

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/util"
	"gorm.io/gorm"
)

func initSqlite() *gorm.DB {
	util.CreateDirNotExists("data/db")
	dsn := fmt.Sprintf("data/db/arw.db")
	sdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ERROR.Println(fmt.Sprintf("database connected err:%v", err))
	}
	if err != nil {
		logger.ERROR.Println(fmt.Sprintf("database connected err:%v", err))
	}
	return sdb
}
