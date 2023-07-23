package router

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/request"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func renderIndex(c *gin.Context) {
	r.LoadHTMLFiles("resource/index.html")
	//r.StaticFile("favicon.ico", "resource/html/favicon.ico")

	c.HTML(http.StatusOK, "index.html", "")
}

func getTvInfo(c *gin.Context) {
	idTemp := c.Param("id")
	idTemp = strings.Split(idTemp, "-")[0]
	language := c.Param("language")
	if language == "" {
		language = "zh-CN"
	}
	re := regexp.MustCompile("[0-9]+")
	idString := re.FindAllString(idTemp, -1)
	id, err := strconv.Atoi(strings.Join(idString, ""))
	if err != nil {
		logger.ERROR.Println(err.Error())
		fail(c, err.Error())
	}
	client := request.NewClient()
	options := map[string]string{
		"language":           language,
		"append_to_response": "credits",
	}
	tvInfo, err := client.GetTVDetails(id, options)
	if err != nil {
		fail(c, err.Error())
	}
	backthread.NewTVNfo(tvInfo)
	marshal, err := json.Marshal(&tvInfo)
	if err != nil {
		return
	}
	successWithData(c, string(marshal))
}
func getPaths(c *gin.Context) {
	data := map[string]string{
		"tv_path":    config.Conf.TVPath,
		"movie_path": config.Conf.MoviePath,
	}
	successWithData(c, data)
}
func changeTVPath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeTVPath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		fail(c, err.Error())
	} else {
		success(c)
	}

}
func changeMoviePath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeMoviePath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		fail(c, err.Error())
	} else {
		success(c)
	}

}
func changeRawPath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeRawPath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		fail(c, err.Error())
	} else {
		success(c)
	}
}
func addRule(c *gin.Context) {
	rawJson, _ := c.GetRawData()
	result := gjson.ParseBytes(rawJson)
	id := result.Get("vid").String()
	provider := result.Get("provider").String()
	fileExtract := result.Get("file_extract").String()
	season := result.Get("season").String()
	language := result.Get("language").String()
	episodeExtract := result.Get("episode_extract").String()
	episodePosition := result.Get("episode_position").String()
	episodeOffset := result.Get("episode_offset").String()

	// default value
	if provider == "" {
		provider = "TMDB"
	}
	if language == "" {
		language = "zh-CN"
	}
	if episodeExtract == "" {
		episodeExtract = `\d+`
	}
	if episodePosition == "" {
		episodePosition = "1"
	}
	if episodeOffset == "" {
		episodeOffset = "0"
	}
	err := backthread.NewRule(id, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset)
	if err != nil {
		fail(c, err.Error())
		logger.ERROR.Println(err.Error())
	} else {
		success(c)
	}

}
func getRule(c *gin.Context) {
	rules := backthread.GetRule()
	data := make([]Rule, 0)
	for _, rule := range rules {
		if !rule.Hidden {
			data = append(data, ExposeRule(rule))
		}
	}

	successWithData(c, data)
}
func hiddenRule(c *gin.Context) {
	id := c.Query("id")
	did, err := strconv.Atoi(id)
	if err != nil {
		logger.ERROR.Println(errors.New("id is not int"))
	}
	err = backthread.HideRule(did)
	if err != nil {
		fail(c, err.Error())
	} else {
		success(c)
	}
}
func updateRule(c *gin.Context) {

}

func getRaw(c *gin.Context) {
	files := backthread.GetFiles()
	successWithData(c, files)
}
