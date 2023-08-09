package backthread

import (
	"errors"
	"github/Yoak3n/anime-repository-web/backend/database"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/package/logger"
	"regexp"
	"strconv"
)

func init() {
	SyncCache()
}

// SyncCache 同步缓存
func SyncCache() {
	db := database.GetDB()
	rules := db.Find(&model.Rules{})
	rows, err := rules.Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		rule := new(model.Rules)
		err = rows.Scan(&rule)
		if err != nil {
			logger.ERROR.Println(err)
			return
		}
		cacheRule := new(model.Rule)
		cacheRule.ID = rule.ID
		cacheRule.VID = rule.VID
		cacheRule.Provider = rule.Provider
		cacheRule.FileExtractReg = regexp.MustCompile(rule.FileExtractReg)
		cacheRule.Name = rule.Name
		cacheRule.Season = rule.Season
		cacheRule.Language = rule.Language
		cacheRule.EpisodeExtractReg = regexp.MustCompile(rule.EpisodeExtractReg)
		cacheRule.EpisodePosition = rule.EpisodePosition
		cacheRule.EpisodeOffset = rule.EpisodeOffset
		cacheRule.Hidden = false
		cache.Rules = append(cache.Rules, cacheRule)
	}
}

func NewRule(vid, name, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset string) error {
	// database
	ruleModel := new(model.Rules)
	ruleModel.VID = vid
	ruleModel.Provider = provider
	ruleModel.FileExtractReg = fileExtract
	ruleModel.Name = name
	s, err := strconv.Atoi(season)
	if err != nil {
		return errors.New("season must be integer")
	}
	ruleModel.Season = s
	ruleModel.Language = language
	ruleModel.EpisodeExtractReg = episodeExtract
	p, err := strconv.Atoi(episodePosition)
	if err != nil {
		return errors.New("episode position must be integer")
	}
	ruleModel.EpisodePosition = p
	o, err := strconv.Atoi(episodeOffset)
	if err != nil {
		return errors.New("episode offset must be integer")
	}
	ruleModel.EpisodeOffset = o
	database.DB.Create(&ruleModel)

	// cache
	ruleCache := new(model.Rule)
	ruleCache.ID = ruleModel.ID
	ruleCache.VID = ruleModel.VID
	ruleCache.Provider = ruleModel.Provider
	ruleCache.FileExtractReg = regexp.MustCompile(ruleModel.FileExtractReg)
	ruleCache.Name = ruleModel.Name
	ruleCache.Season = ruleModel.Season
	ruleCache.Language = ruleModel.Language
	ruleCache.EpisodeExtractReg = regexp.MustCompile(ruleModel.EpisodeExtractReg)
	ruleCache.EpisodePosition = ruleModel.EpisodePosition
	ruleCache.EpisodeOffset = ruleModel.EpisodeOffset
	ruleCache.Hidden = false
	cache.Rules = append(cache.Rules, ruleCache)
	err = Scan()
	if err != nil {
		return err
	}
	return nil
}

func ModifyRule(id uint, vid, name, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset string) error {
	// database
	ruleModel := new(model.Rules)
	ruleModel.ID = id
	ruleModel.VID = vid
	ruleModel.Name = name
	ruleModel.Provider = provider
	ruleModel.FileExtractReg = fileExtract
	ruleModel.Name = ""
	s, err := strconv.Atoi(season)
	if err != nil {
		return errors.New("season must be integer")
	}
	ruleModel.Season = s
	ruleModel.Language = language
	ruleModel.EpisodeExtractReg = episodeExtract
	o, err := strconv.Atoi(episodeOffset)
	if err != nil {
		return errors.New("episode offset must be integer")
	}
	ruleModel.EpisodeOffset = o
	p, err := strconv.Atoi(episodePosition)
	if err != nil {
		return errors.New("episode position must be integer")
	}
	ruleModel.EpisodePosition = p
	// cache
	var rule *model.Rule
	for _, r := range cache.Rules {
		if r.ID == id {
			rule = r
		}
	}
	rule.VID = vid
	rule.Provider = provider
	rule.FileExtractReg = regexp.MustCompile(fileExtract)
	rule.Season = s
	rule.Language = language
	rule.EpisodeExtractReg = regexp.MustCompile(episodeExtract)
	rule.EpisodePosition = p
	rule.EpisodeOffset = o
	rule.Hidden = false
	err = Scan()
	if err != nil {
		return err
	}
	return nil
}

func GetRule() []*model.Rule {
	rules := make([]*model.Rule, 0)
	for _, rule := range cache.Rules {
		rules = append(rules, rule)
	}
	return rules
}

func HideRule(id uint) {
	db := database.DB
	db.Where("id = ?", id).Delete(&model.Rules{})

	// 但愿这个协程不会有什么问题
	go func() {
		for _, rule := range cache.Rules {
			if id == rule.ID {
				rule.Hidden = true
			}
		}
	}()
}
