package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/request"
	"github/Yoak3n/anime-repository-web/package/response"
	"regexp"
	"strconv"
	"strings"
)

func GetTvInfo(c *gin.Context) {
	idTemp := c.Param("id")
	idTemp = strings.Split(idTemp, "-")[0]
	language := c.Param("language")
	if language == "" {
		language = "zh"
	}
	re := regexp.MustCompile("[0-9]+")
	idString := re.FindAllString(idTemp, -1)
	id, err := strconv.Atoi(strings.Join(idString, ""))
	if err != nil {
		logger.ERROR.Println(err.Error())
		response.Fail(c, err.Error())
	}
	client := request.GetClient()
	options := map[string]string{
		"language":           language,
		"append_to_response": "credits",
	}
	tvInfo, err := client.GetTVDetails(id, options)
	if err != nil {
		response.Fail(c, err.Error())
	}
	err = backthread.NewTVNfo(tvInfo)
	if err != nil {
		logger.ERROR.Println(err.Error())
		return
	}
	marshal, err := json.Marshal(&tvInfo)
	if err != nil {
		logger.ERROR.Println(err.Error())
		return
	}
	response.SuccessWithData(c, string(marshal))
}

func GetRaw(c *gin.Context) {
	files := backthread.GetFiles()
	response.SuccessWithData(c, files)
}
func GetRawTVName(c *gin.Context) {
	vid := c.Param("vid")
	language := c.Query("language")
	if language == "" {
		language = "zh"
	}
	name, err := request.GetVideoTVName(vid, language)
	if err != nil {
		response.Fail(c, err.Error())
	} else {
		response.SuccessWithData(c, name)
	}
}
