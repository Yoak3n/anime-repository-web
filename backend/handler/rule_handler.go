package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/response"
	"strconv"
)

func AddRule(c *gin.Context) {
	rawJson, _ := c.GetRawData()
	result := gjson.ParseBytes(rawJson)
	id := result.Get("vid").String()
	name := result.Get("name").String()
	provider := result.Get("provider").String()
	fileExtract := result.Get("file_extract_reg").String()
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
		language = "zh"
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
	err := backthread.NewRule(id, name, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset)
	if err != nil {
		logger.ERROR.Println(err.Error())
		response.Fail(c, err.Error())
	} else {
		response.Success(c)
	}

}
func GetRule(c *gin.Context) {
	rules := backthread.GetRule()
	data := make([]response.Rule, 0)
	for _, rule := range rules {
		if !rule.Hidden {
			data = append(data, response.ExposeRule(rule))
		}
	}

	response.SuccessWithData(c, data)
}
func HiddenRule(c *gin.Context) {
	id := c.Query("id")
	did, err := strconv.Atoi(id)
	if err != nil {
		logger.ERROR.Println(errors.New("id is not int"))
	}
	backthread.HideRule(uint(did))
	response.Success(c)

}
func UpdateRule(c *gin.Context) {
	temp := c.Param("id")
	id, err := strconv.Atoi(temp)
	response.Fail(c, err.Error())
	rawJson, _ := c.GetRawData()
	result := gjson.ParseBytes(rawJson)
	vid := result.Get("vid").String()
	name := result.Get("name").String()
	provider := result.Get("provider").String()
	fileExtract := result.Get("file_extract").String()
	season := result.Get("season").String()
	language := result.Get("language").String()
	episodeExtract := result.Get("episode_extract").String()
	episodePosition := result.Get("episode_position").String()
	episodeOffset := result.Get("episode_offset").String()
	err = backthread.ModifyRule(uint(id), vid, name, provider, fileExtract, season, language, episodeExtract, episodePosition, episodeOffset)
	if err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c)
	}
}
