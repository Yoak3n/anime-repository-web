package backthread

import (
	"errors"
	"fmt"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/request"
	"github/Yoak3n/anime-repository-web/package/util"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cache struct {
	TvFiles    map[string]*model.TVItem
	MatchRules map[string]*model.Rule
	Recognized map[string]bool
	Rules      []*model.Rule
}

// create a cache for disposal
var cache *Cache

func init() {
	cache = &Cache{
		TvFiles:    make(map[string]*model.TVItem),
		MatchRules: make(map[string]*model.Rule),
		Rules:      make([]*model.Rule, 0),
		Recognized: make(map[string]bool),
	}
}

// Dispose cache to make files match rules
func Dispose(filePath string) {
	logger.DEBUG.Println("start dispose:", filePath)
	for path, _ := range cache.TvFiles {
		if path == filePath {
			return
		}
	}
	i := model.NewTVItem(filePath)
	cache.TvFiles[filePath] = i
	cache.Recognized[filePath] = false

}

func MatchRules() {
	logger.DEBUG.Println("start match rules")
	for filePath, _ := range cache.TvFiles {
		for _, rule := range cache.Rules {
			if rule.FileExtractReg.Match([]byte(filePath)) && !rule.Hidden && cache.Recognized[filePath] {
				logger.DEBUG.Println("匹配成功")
				cache.MatchRules[filePath] = rule
			}
		}
	}

}

func WorkOnMatchedCache() error {

	logger.DEBUG.Println("start work on matched cache")
	for filePath, rule := range cache.MatchRules {
		// 过滤已识别文件，不过之后已识别的文件已经不在该文件夹
		if !cache.Recognized[filePath] {
			allDirectory := fmt.Sprintf("%s/%s/Season 0%d", config.Conf.TVPath, rule.Name, rule.Season)
			// 开始刮削图片
			logger.INFO.Println("start cutting image:", filePath)
			tc := request.GetClient()
			options := map[string]string{
				"language":           rule.Language,
				"append_to_response": "credits,images",
			}
			re := regexp.MustCompile("[0-9]+")
			idString := strings.SplitN(rule.VID, "-", 1)
			idString = re.FindAllString(idString[0], -1)
			id, err := strconv.Atoi(strings.Join(idString, ""))

			for {
				tvInfo, e := tc.GetTVDetails(id, options)
				if e != nil {
					logger.ERROR.Println(e)
				}
				e = NewTVNfo(tvInfo)
				if e != nil {
					continue
				} else {
					break
				}
			}

			util.CreateDirNotExists(allDirectory)
			fileName := cache.TvFiles[filePath].Name
			episode, err := strconv.Atoi(rule.EpisodeExtractReg.FindAllString(fileName, -1)[rule.EpisodePosition-1])
			episode = episode + rule.EpisodeOffset
			if err != nil {
				return errors.New("can't get episode number")
			}
			newPath := fmt.Sprintf("%s/%s - S0%dE0%d%s", allDirectory, rule.Name, rule.Season, episode, cache.TvFiles[filePath].Suffix)
			err = os.Rename(filePath, newPath)
			if err != nil {
				return err
			}
			err = os.Link(newPath, filePath)
			if err != nil {
				return err
			}

			tvEpisodeInfo, err := request.GetTVEpisode(id, rule.Season, episode)
			if err != nil {
				return err
			}
			go func() {
				err = handleTVEpisodeNfo(tvEpisodeInfo, allDirectory, rule.Name, rule.Season)
				if err != nil {
					logger.ERROR.Println(err)
				}
			}()

			cache.Recognized[filePath] = true
			logger.DEBUG.Println(fmt.Sprintf("matched cache: %s", filePath))
		}
	}
	return nil
}
