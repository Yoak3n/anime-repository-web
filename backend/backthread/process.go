package backthread

import (
	"errors"
	"github/Yoak3n/anime-repository-web/backend/item"
	"github/Yoak3n/anime-repository-web/backend/model"
	"regexp"
	"strconv"
)

type Cache struct {
	TvFiles    map[string]*item.TVItem
	MatchRules map[string]*model.Rule
	Rules      []*model.Rule
}

// create a cache for disposal
var cache *Cache

func init() {
	cache = &Cache{
		TvFiles:    make(map[string]*item.TVItem),
		MatchRules: make(map[string]*model.Rule),
		Rules:      make([]*model.Rule, 0),
	}
}
func Dispose(filePath string) {
	for path, _ := range cache.TvFiles {
		if path == filePath {
			return
		}
	}
	for _, rule := range cache.Rules {
		if rule.FileExtractReg.Match([]byte(filePath)) {
			cache.MatchRules[filePath] = rule
		}
	}
	i := item.NewTVItem(filePath)
	cache.TvFiles[filePath] = i
}

func NewRule(id, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset string) error {
	rule := new(model.Rule)

	rule.VID = id
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
	rule.ID = len(cache.Rules) + 1
	// 后续会使用数据库，暂时就这样不考虑重复
	cache.Rules = append(cache.Rules, rule)
	return nil
}

func GetRule(vid string) *model.Rule {
	for _, rule := range cache.Rules {
		if vid == rule.VID {
			return rule
		}
	}
	return nil
}
