package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/response"
	"net/http"
)

func GetConfig(c *gin.Context) {
	response.SuccessWithData(c, response.ExposeConfig(config.Conf))
}
func ChangeConfig(c *gin.Context) {
	rawJson, err := c.GetRawData()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, "Get json error", nil)
	} else {
		result := gjson.ParseBytes(rawJson)
		key := result.Get("key").String()
		useProxy := result.Get("use_proxy").Bool()
		proxyUrl := result.Get("proxy_url").String()
		rawPath := result.Get("raw_path").String()
		tvPath := result.Get("tv_path").String()
		moviePath := result.Get("movie_path").String()
		delay := result.Get("delay").Int()
		err = config.ChangeAllConfig(key, useProxy, proxyUrl, rawPath, tvPath, moviePath, delay)
		if err != nil {
			response.Response(c, http.StatusInternalServerError, "Change config error", nil)
		} else {
			_ = backthread.Scan()
			response.Success(c)
		}
	}
}
