package database

import (
	"database/sql"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/package/logger"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Conn *sql.DB

func init() {
	DB = initSqlite()
	Conn, _ = DB.DB()
	err := DB.AutoMigrate(&model.RuleModel{})
	if err != nil {
		logger.ERROR.Println("migrate error:", err)
		return
	}
	defer Conn.Close()
}
