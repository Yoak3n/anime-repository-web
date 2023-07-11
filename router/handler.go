package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/request"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func renderIndex(c *gin.Context) {
	r.LoadHTMLFiles("resource/html/index.html")
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
		fail(c, err.Error())
	}
	success(c)
}

func changeMoviePath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeMoviePath(path)
	if err != nil {
		fail(c, err.Error())
	}
	success(c)
}
func changeRawPath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeRawPath(path)
	if err != nil {
		fail(c, err.Error())
	}
	success(c)
}
