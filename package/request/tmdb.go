package request

import (
	"errors"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/config"
	"io"
	"log"
	"net/http"
	"time"
)

func NewClient() *tmdb.Client {
	TC, err := tmdb.Init(config.Conf.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	customClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
			Proxy:           http.ProxyURL(config.Conf.Proxy),
		},
	}
	TC.SetClientConfig(customClient)
	log.Println("TMDB Client initialized")
	return TC
}


func GetVideoName(id string) (string, error) {
		return "",nil
}
func NewToken() (string, error) {
	res, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/authentication/token/new?api_key=%s", config.Conf.ApiKey))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	jsonResult := gjson.ParseBytes(result)
	if jsonResult.Get("success").Bool() == false {
		return "", errors.New("failed to get token")
	}
	return jsonResult.Get("request_token").String(), nil
}
