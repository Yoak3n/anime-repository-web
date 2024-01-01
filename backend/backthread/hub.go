package backthread

import (
	"errors"
	"fmt"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/request"
	"github/Yoak3n/anime-repository-web/package/util"
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
	for path := range cache.TvFiles {
		if path == filePath {
			return
		}
	}
	i := model.NewTVItem(filePath)
	cache.TvFiles[filePath] = i

}

func MatchRules() {
	logger.DEBUG.Println("start match rules")
	for filePath := range cache.TvFiles {
		for _, rule := range cache.Rules {
			if !cache.Recognized[filePath] {
				if rule.FileExtractReg.Match([]byte(filePath)) && !rule.Hidden && !cache.Recognized[filePath] {
					logger.DEBUG.Println("匹配成功")
					cache.MatchRules[filePath] = rule
				}
			} else {
				logger.INFO.Println("文件已识别")
			}

		}
	}

}

func WorkOnMatchedCache() error {
	logger.INFO.Println("处理文件")
	logger.DEBUG.Println("start work on matched cache")
	for filePath, rule := range cache.MatchRules {
		// 过滤已识别文件，不过之后已识别的文件已经不在该文件夹
		if !cache.Recognized[filePath] {
			allDirectory := fmt.Sprintf("%s/%s/Season %d", config.Conf.TVPath, rule.Name, rule.Season)
			util.CreateDirNotExists(allDirectory)
			logger.INFO.Println("Create dir:", allDirectory)

			// 开始刮削图片
			logger.INFO.Println("start cutting image:", filePath, allDirectory)
			tc := request.GetClient()
			options := map[string]string{
				"language":           rule.Language,
				"append_to_response": "credits,images",
			}
			re := regexp.MustCompile("[0-9]+")
			idString := strings.SplitN(rule.VID, "-", 1)
			idString = re.FindAllString(idString[0], -1)
			id, err := strconv.Atoi(strings.Join(idString, ""))

			// 发送请求
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

			fileName := cache.TvFiles[filePath].Name
			episode, err := strconv.Atoi(rule.EpisodeExtractReg.FindAllString(fileName, -1)[rule.EpisodePosition-1])
			episode = episode + rule.EpisodeOffset
			if err != nil {
				return errors.New("can't get episode number")
			}
			newPath := fmt.Sprintf("%s/%s - S%dE%d%s", allDirectory, rule.Name, rule.Season, episode, cache.TvFiles[filePath].Suffix)
			err = util.MoveVideoFile(filePath, newPath)
			if err != nil {
				logger.DEBUG.Println(err)
				return err
			}

			tvEpisodeInfo, err := request.GetTVEpisode(id, rule.Season, episode, rule.Language)
			if err != nil {
				return err
			}
			go func() {
				err = handleTVEpisodeNfo(tvEpisodeInfo, allDirectory, rule.Name, rule.Season)
				if err != nil {
					logger.DEBUG.Println(err)
				}
			}()

			cache.Recognized[filePath] = true
			logger.DEBUG.Println(fmt.Sprintf("cache matched: %s", filePath))
		}
	}
	return nil
}
