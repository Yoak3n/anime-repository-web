package util

import (
	"bufio"
	"encoding/xml"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/model"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	themovidb = "themovidb"
)

func NewTVNfo(tv *tmdb.TVDetails) {
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
	// clearlogo
	thumbs = append(thumbs, model.Thumb{Aspect: "clearlogo", Value: makeImagePath(tv.Networks[0].LogoPath)})
	// genSeasonPoster
	// for collection
	seasonPoster := make([]string, 0)
	seasonNumber := make([]int, 0)
	for i := 0; i < len(tv.Seasons); i++ {
		poster := makeImagePath(tv.Seasons[i].PosterPath)
		number := tv.Seasons[i].SeasonNumber

		seasonPoster = append(seasonPoster, poster)
		seasonNumber = append(seasonNumber, number)

		thumbs = append(thumbs, model.Thumb{Aspect: "poster", Value: poster, Season: number, Type: "season"})
	}
	thumbs = append(thumbs, model.Thumb{Aspect: "poster", Value: makeImagePath(tv.PosterPath)})
	// genActors
	actors := make([]model.Actor, 0)
	for _, actor := range tv.TVCreditsAppend.Credits.Cast {
		actors = append(actors, model.Actor{Name: actor.Name, Role: actor.Character, Order: actor.Order, Thumb: model.Thumb{Value: makeImagePath(actor.ProfilePath)}})
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
			model.Thumb{Value: makeImagePath(tv.BackdropPath)},
		},
	}
	nfoData, err := xml.Marshal(&tvNfo)
	if err != nil {
		log.Println("err")
	}
	go writeTVShowNfo(nfoData, tv.Name)
	go collectImages(tv.Networks[0].LogoPath, tv.PosterPath, tv.BackdropPath, seasonPoster, seasonNumber)

}

func writeTVShowNfo(nfo []byte, title string) {
	HEADER := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>`)
	data := append(HEADER, nfo...)
	err := os.WriteFile(fmt.Sprintf("/%s/%s/tvshow.nfo", config.Conf.TVPath, title), data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func collectImages(logo string, poster string, fanart string, seasonsPosters []string, seasonNumber []int) {
	seasonsImages := make([]map[string]string, 0)
	for i := 0; i < len(seasonsPosters); i++ {
		seasonsImages = append(seasonsImages, map[string]string{
			fmt.Sprintf("season0%d-poster.jpg", seasonNumber[i]): seasonsPosters[i],
		})
	}
	collection := model.TVCollection{
		ClearLogo: makeImagePath(logo),
		Poster:    makeImagePath(poster),
		Fanart:    makeImagePath(fanart),
		Seasons:   seasonsImages,
	}
	downloadImages(&collection)
}

func downloadImages(c *model.TVCollection) {
	// 同步or异步？——是否会被风控

}
func downloadImage(path string, imgUrl string) error {
	res, err := http.Get(imgUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 4096)

	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer, reader)
	if err != nil {
		log.Println("写入文件错误：", err)
		return err
	}
	fmt.Printf("Total length: %d", written)
	return nil
}
