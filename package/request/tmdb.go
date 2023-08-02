package request

import (
	"errors"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"io"
	"net/http"
	"time"
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
	uri := fmt.Sprintf("https://api.themoviedb.org/3/tv/%s?language=%s&api_key=%s", vid, language, config.Conf.ApiKey)
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

func GetTVEpisode(id int, season int, episode int) (*model.TVEpisode, error) {
	imagePrefix := "https://image.tmdb.org/t/p/original"
	uri := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d/season/%d/episode/%d?api_key=%s&append_to_response=credits,images", id, season, episode, config.Conf.ApiKey)

	client := GenRequestClient()
	res, err := client.Get(uri)
	if res.StatusCode != 200 || err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	result := gjson.ParseBytes(buf)
	episodeInfo := new(model.TVEpisode)
	// Episode
	episodeInfo.Episode = result.Get("episode_number").Int()
	// Season
	episodeInfo.Season = result.Get("season_number").Int()
	// Plot
	episodeInfo.Plot = result.Get("overview").String()
	// Rating
	episodeInfo.Ratings.Rating = make([]model.Rating, 1)
	episodeInfo.Ratings.Rating[0].Name = "themoviedb"
	episodeInfo.Ratings.Rating[0].Max = 10
	episodeInfo.Ratings.Rating[0].Default = true
	episodeInfo.Ratings.Rating[0].Value = float32(result.Get("vote_average").Float())
	episodeInfo.Ratings.Rating[0].Votes = result.Get("vote_count").Int()
	// Thumb
	episodeInfo.Thumb.Aspect = "thumb"
	episodeInfo.Thumb.Value = imagePrefix + result.Get("still_path").String()
	// UniqueID
	episodeInfo.UniqueID.Value = result.Get("id").String()
	episodeInfo.UniqueID.Type = "tmdb"
	episodeInfo.UniqueID.Default = true
	// Aired
	episodeInfo.Aired = result.Get("air_date").String()
	// Actor
	actors := result.Get("credits.cast").Array()
	for i := 0; i < len(actors); i++ {
		actor := model.Actor{
			Name:  actors[i].Get("name").String(),
			Role:  actors[i].Get("character").String(),
			Order: int(actors[i].Get("order").Int()),
			Thumb: model.Thumb{
				Value: imagePrefix + actors[i].Get("profile_path").String(),
			},
		}
		episodeInfo.Actor = append(episodeInfo.Actor, actor)
	}
	// DateAdded
	episodeInfo.DateAdded = time.Now().Format(time.DateTime)
	return episodeInfo, nil
}
