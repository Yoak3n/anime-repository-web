package backthread

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"github/Yoak3n/anime-repository-web/package/util"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	tmdb "github.com/cyruzin/golang-tmdb"
)

const (
	themovidb = "themovidb"
)

func NewTVNfo(tv *tmdb.TVDetails) error {
	if tv == nil {
		return errors.New("tv is nil")
	}
	// genRatings
	ratings := model.Rating{
		Name:    themovidb,
		Max:     10,
		Default: true,
		Value:   tv.VoteAverage,
		Votes:   tv.VoteCount,
	}
	// genStudios
	studios := make([]string, 0)
	for _, company := range tv.ProductionCompanies {
		studios = append(studios, company.Name)
	}
	// genGenres
	genres := make([]string, 0)
	for _, genre := range tv.Genres {
		genres = append(genres, genre.Name)
	}
	// genThumbs
	thumbs := make([]model.Thumb, 0)
	if len(tv.Networks) > 0 {
		// genClearlogo
		thumbs = append(thumbs, model.Thumb{Aspect: "clearlogo", Value: util.MakeImagePath(tv.Networks[0].LogoPath)})
	}
	// genSeasonPoster
	// for collection
	seasonPoster := make([]string, 0)
	seasonNumber := make([]int, 0)
	for i := 0; i < len(tv.Seasons); i++ {
		poster := util.MakeImagePath(tv.Seasons[i].PosterPath)
		number := tv.Seasons[i].SeasonNumber

		seasonPoster = append(seasonPoster, poster)
		seasonNumber = append(seasonNumber, number)

		thumbs = append(thumbs, model.Thumb{Aspect: "poster", Value: poster, Season: number, Type: "season"})
	}
	thumbs = append(thumbs, model.Thumb{Aspect: "poster", Value: util.MakeImagePath(tv.PosterPath)})
	// genActors
	actors := make([]model.Actor, 0)
	for _, actor := range tv.TVCreditsAppend.Credits.Cast {
		actors = append(actors, model.Actor{Name: actor.Name, Role: actor.Character, Order: actor.Order, Thumb: model.Thumb{Value: util.MakeImagePath(actor.ProfilePath)}})
	}
	tvNfo := model.TVShow{
		Title:         tv.Name,
		OriginalTitle: tv.OriginalName,
		Premiered:     tv.FirstAirDate,
		Ratings: model.Ratings{
			Rating: []model.Rating{ratings},
		},
		Plot:   tv.Overview,
		Studio: studios,
		Genre:  genres,
		Thumb:  thumbs,
		Fanart: []model.Thumb{
			{
				Value: util.MakeImagePath(tv.BackdropPath),
			},
		},
		UniqueID: model.UniqueID{
			Value:   strconv.FormatInt(tv.ID, 10),
			Type:    "tmdb",
			Default: true,
		},
		Actor: actors,
	}
	nfoData, err := xml.Marshal(&tvNfo)
	if err != nil {
		logger.ERROR.Println(err)
	}

	logoPath := tv.Networks[0].LogoPath
	posterPath := tv.PosterPath
	backdropPath := tv.BackdropPath

	if len(tv.Networks) > 0 {
		go writeTVShowNfo(nfoData, tv.Name)
		go collectImages(tv.Name, util.MakeImagePath(logoPath), util.MakeImagePath(posterPath), util.MakeImagePath(backdropPath), seasonPoster, seasonNumber)
	} else {
		name := tv.Name
		go writeTVShowNfo(nfoData, name)
		go collectImages(name, "", posterPath, backdropPath, seasonPoster, seasonNumber)
	}

	return nil
}

func writeTVShowNfo(nfo []byte, title string) {
	HEADER := []byte(`<?xml version="0.0" encoding="UTF-8" standalone="yes" ?>`)
	data := append(HEADER, nfo...)
	err := os.WriteFile(fmt.Sprintf("%s/%s/tvshow.nfo", config.Conf.TVPath, title), data, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func collectImages(name string, logo string, poster string, fanart string, seasonsPosters []string, seasonNumber []int) {
	seasonsImages := make(map[string]string)
	for i := 0; i < len(seasonsPosters); i++ {
		seasonsImages[fmt.Sprintf("Season%d-poster.jpg", seasonNumber[i])] = seasonsPosters[i]
	}
	collection := new(model.TVCollection)
	if logo != "" {
		collection = &model.TVCollection{
			ClearLogo: util.MakeImagePath(logo),
			Poster:    util.MakeImagePath(poster),
			Fanart:    util.MakeImagePath(fanart),
			Seasons:   seasonsImages,
		}
	} else {
		collection = &model.TVCollection{
			Poster:  util.MakeImagePath(poster),
			Fanart:  util.MakeImagePath(fanart),
			Seasons: seasonsImages,
		}
	}

	downloadImages(name, collection)
}

func downloadImages(name string, c *model.TVCollection) {
	// 同步or异步？---是否会被风控
	wg := &sync.WaitGroup{}
	wg.Add(3 + len(c.Seasons))

	tvDir := fmt.Sprintf("%s/%s/", config.Conf.TVPath, name)
	clearLogo := c.ClearLogo
	poster := c.Poster
	fanart := c.Fanart

	go func() {
		defer wg.Done()
		err := downloadImage(tvDir+"clearlogo.png", clearLogo)
		logger.INFO.Println("downloading clearlogo")
		if err != nil {
			log.Println(name, "clearlogo download error:", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := downloadImage(tvDir+"poster.jpg", poster)
		logger.INFO.Println("downloading poster")
		if err != nil {
			log.Println(name, "poster download error:", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := downloadImage(tvDir+"fanart.jpg", fanart)
		logger.INFO.Println("downloading fanart")
		if err != nil {
			logger.INFO.Println(name, "fanart download error:", err)
		}
	}()
	for k, v := range c.Seasons {
		key := k
		value := v
		go func() {
			defer wg.Done()
			err := downloadImage(fmt.Sprintf("%s/%s/%s", config.Conf.TVPath, name, key), value)
			if err != nil {
				logger.INFO.Println(name, "season poster download error:", err)
			}
		}()
	}
	wg.Wait()
}
func downloadImage(path string, imgUrl string) error {
	if _, err := os.Stat(path); os.IsExist(err) {
		return err
	}
	res, err := http.Get(imgUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 4096)

	file, err := os.Create(path)
	if err != nil {
		logger.ERROR.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer, reader)
	if err != nil {
		log.Println("write file error: ", err)
		return err
	}
	fmt.Printf("Total length: %d\n", written)
	return nil
}

func writeTVEpisodeNfo(data []byte, tvPath string, name string) error {
	HEADER := []byte(`<?xml version="0.0" encoding="UTF-8" standalone="yes" ?>`)
	nfoData := append(HEADER, data...)
	err := os.WriteFile(fmt.Sprintf("%s/%s.nfo", tvPath, name), nfoData, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func handleTVEpisodeNfo(nfo *model.TVEpisode, newPath string, name string, season int) error {
	tvEpisodeData, err := xml.Marshal(nfo)
	if err != nil {
		return err
	}
	episodeName := fmt.Sprintf("%s - S%dE%d - %s", name, nfo.Season, nfo.Episode, nfo.Title)
	err = writeTVEpisodeNfo(tvEpisodeData, newPath, episodeName)
	if err != nil {
		return err
	}
	episodeImg := fmt.Sprintf("%s/%s-thumb.jpg", newPath, episodeName)
	thumbUrl := util.MakeImagePath(nfo.Thumb.Value)
	err = downloadImage(episodeImg, thumbUrl)
	if err != nil {
		return err
	}
	return nil
}
