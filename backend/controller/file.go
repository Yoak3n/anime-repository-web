package controller

import (
	"github/Yoak3n/anime-repository-web/backend/database"
	"github/Yoak3n/anime-repository-web/backend/model"
)

func AppendNewFile(fp model.VideoFiles) {
	db := database.GetDB()
	if db.First(&fp).Where("recognized=?", true).RowsAffected != 0 {
		return
	}
	if fp.Recognized {
		return
	}
	db.Create(&fp)
}
