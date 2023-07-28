package handler

import (
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/response"
)

func GetPaths(c *gin.Context) {
	data := map[string]string{
		"tv_path":    config.Conf.TVPath,
		"movie_path": config.Conf.MoviePath,
	}
	response.SuccessWithData(c, data)
}
func ChangeTVPath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeTVPath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		response.Fail(c, err.Error())
	} else {
		response.Success(c)
	}

}
func ChangeMoviePath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeMoviePath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		response.Fail(c, err.Error())
	} else {
		response.Success(c)
	}

}
func ChangeRawPath(c *gin.Context) {
	path := c.Param("path")
	err := config.ChangeRawPath(path)
	if err != nil {
		logger.ERROR.Println(err.Error())
		response.Fail(c, err.Error())
	} else {
		response.Success(c)
	}
}
