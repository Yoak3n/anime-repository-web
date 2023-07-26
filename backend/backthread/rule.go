package backthread

import (
	"errors"
	"github/Yoak3n/anime-repository-web/backend/model"
	"regexp"
	"strconv"
)

func NewRule(vid, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset string) error {
	rule := new(model.Rule)

	rule.VID = vid
	rule.Provider = provider
	rule.FileExtractReg = regexp.MustCompile(fileExtract)
	s, err := strconv.Atoi(season)
	if err != nil {
		return errors.New("season must be integer")
	}
	rule.Season = s
	rule.Language = language
	rule.EpisodeExtractReg = regexp.MustCompile(episodeExtract)
	p, err := strconv.Atoi(episodePosition)
	if err != nil {
		return errors.New("episode position must be integer")
	}
	rule.EpisodePosition = p
	o, err := strconv.Atoi(episodeOffset)
	if err != nil {
		return errors.New("episode offset must be integer")
	}
	rule.EpisodeOffset = o
	rule.Hidden = false
	rule.ID = uint(len(cache.Rules) + 1)
	// 后续可能会使用数据库实现数据持久，暂时就这样不考虑重复
	cache.Rules = append(cache.Rules, rule)

	return nil
}

func ModifyRule(id int64, vid, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset string) error {
	rule := *cache.Rules[id-1]
	rule.VID = vid
	rule.Provider = provider
	rule.FileExtractReg = regexp.MustCompile(fileExtract)
	s, err := strconv.Atoi(season)
	if err != nil {
		return errors.New("season must be integer")
	}
	rule.Season = s
	rule.Language = language
	rule.EpisodeExtractReg = regexp.MustCompile(episodeExtract)
	p, err := strconv.Atoi(episodePosition)
	if err != nil {
		return errors.New("episode position must be integer")
	}
	rule.EpisodePosition = p
	o, err := strconv.Atoi(episodeOffset)
	if err != nil {
		return errors.New("episode offset must be integer")
	}
	rule.EpisodeOffset = o
	rule.Hidden = false

	return nil
}

func GetRule() []*model.Rule {
	rules := make([]*model.Rule, 0)
	for _, rule := range cache.Rules {
		rules = append(rules, rule)
	}
	return rules
}

func HideRule(id uint) error {
	for _, rule := range cache.Rules {
		if id == rule.ID {
			rule.Hidden = true
			return nil
		}
	}
	return errors.New("rule not found")
}
