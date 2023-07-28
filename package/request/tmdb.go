package request

import (
	"errors"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"io"
	"net/http"
)

var TC *tmdb.Client
var e error

func init() {
	TC, e = tmdb.Init(config.Conf.ApiKey)
	if e != nil {
		logger.ERROR.Panic(e)
	}
	TC.SetClientConfig(GenRequestClient())
	logger.INFO.Println("TMDB Client initialized")
}

func GetClient() *tmdb.Client {
	return TC
}

func GetVideoTVName(vid string, language string) (string, error) {
	uri := fmt.Sprintf("https://api.themoviedb.org/3/tv/%S?language=%s&api_key=%s", vid, language, config.Conf.ApiKey)
	client := GenRequestClient()
	res, err := client.Get(uri)
	if err != nil || res.StatusCode != 200 {
		return "", err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	result := gjson.ParseBytes(buf)
	name := result.Get("name").String()
	return name, nil
}
func NewToken() (string, error) {
	res, e := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/authentication/token/new?api_key=%s", config.Conf.ApiKey))
	if e != nil {
		return "", e
	}
	defer res.Body.Close()
	result, e := io.ReadAll(res.Body)
	if e != nil {
		return "", e
	}
	jsonResult := gjson.ParseBytes(result)
	if jsonResult.Get("success").Bool() == false {
		return "", errors.New("failed to get token")
	}
	return jsonResult.Get("request_token").String(), nil
}

func GetTVEpisode(id int, season int, episode int) (string, error) {
	uri := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d/season/%d/episode/%d?api_key=%s&append_to_response=credits,images", id, season, episode, config.Conf.ApiKey)

	client := GenRequestClient()
	res, err := client.Get(uri)
	if res.StatusCode != 200 || err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	result := gjson.ParseBytes(buf)
	title := result.Get("name").String()
	return title, err
}
