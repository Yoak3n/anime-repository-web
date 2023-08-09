package database

import (
	"database/sql"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/package/logger"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var conn *sql.DB

func init() {
	db = initSqlite()
	conn, _ = db.DB()
	err := db.AutoMigrate(&model.Rules{})
	if err != nil {
		logger.ERROR.Println("table migrate error:", err)
		return
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(time.Hour)

}

func GetDB() *gorm.DB {
	return db
}

func GetConn() *sql.DB {
	return conn
}
