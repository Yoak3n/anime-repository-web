package backthread

import (
	"fmt"
	"github/Yoak3n/anime-repository-web/backend/item"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/request"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cache struct {
	TvFiles    map[string]*item.TVItem
	MatchRules map[string]*model.Rule
	Recognized map[string]bool
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

// Dispose cache to make files match rules
func Dispose(filePath string) {
	for path, _ := range cache.TvFiles {
		if path == filePath {
			return
		}
	}
	for _, rule := range cache.Rules {
		if rule.FileExtractReg.Match([]byte(filePath)) && rule.Hidden == false {
			cache.MatchRules[filePath] = rule
		}
	}
	i := item.NewTVItem(filePath)
	cache.TvFiles[filePath] = i
}

func WorkOnMatchedCache() error {
	for filePath, rule := range cache.MatchRules {
		// 过滤已识别
		if !cache.Recognized[filePath] {
			allDirectory := fmt.Sprintf("%s/%s/season %d", config.Conf.TVPath, rule.Name, rule.Season)
			err := os.MkdirAll(allDirectory, os.ModePerm)
			if err != nil {
				return err
			}
			// 开始刮削图片
			tc := request.GetClient()
			options := map[string]string{
				"language":           rule.Language,
				"append_to_response": "credits",
			}
			re := regexp.MustCompile("[0-9]+")
			idString := strings.SplitN(rule.VID, "-", 1)
			idString = re.FindAllString(idString[0], -1)
			id, err := strconv.Atoi(strings.Join(idString, ""))
			tvInfo, err := tc.GetTVDetails(id, options)
			NewTVNfo(tvInfo)
			newPath := allDirectory + "/" + cache.TvFiles[filePath].Name
			err = os.Rename(filePath, newPath)
			if err != nil {
				return err
			}
			err = os.Link(newPath, filePath)
			if err != nil {
				return err
			}
			cache.Recognized[filePath] = true
		}
	}
	return nil
}
