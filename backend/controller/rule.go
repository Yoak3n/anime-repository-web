package controller

import (
	"github/Yoak3n/anime-repository-web/backend/database"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/package/logger"
)

func GetRules() (rules []model.Rules) {
	db := database.GetDB()
	result := db.Find(&rules) // 获取所有规则
	if result.Error != nil {
		logger.ERROR.Println(result.Error)
		return nil
	}
	return rules

}
