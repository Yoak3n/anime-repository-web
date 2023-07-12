package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"log"
	"net/http"
	"os"
	"time"
)

type TVShow struct {
	XMLName       xml.Name `xml:"tvshow"`
	Title         string   `xml:"title"`
	OriginalTitle string   `xml:"originaltitle"`
	SortTitle     string
	Ratings       Ratings `xml:"ratings"`
	UserRating    string  `xml:",omitempty"`
	Top250        int
	Season        int
	Episode       int
	Plot          string `xml:"plot"`
	Tagline       string
	Thumb         []Thumb `xml:"thumb"`
	Fanart        []Thumb `xml:"fanart"`
	Mpaa          string
	PlayCount     int
	LastPlayed    string
	UniqueID      string `xml:"uniqueid"`
}
type Ratings struct {
	Rating []Rating `xml:"rating"`
}
type Rating struct {
	Name    string  `xml:"name,attr"`
	Max     int     `xml:"max,attr"`
	Default bool    `xml:"default,attr"`
	Value   float64 `xml:"value"`
	Votes   int     `xml:"votes"`
}
type Thumb struct {
	Aspect  string `xml:"aspect,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Season  int    `xml:"season,attr,omitempty"`
	Preview string `xml:"preview,omitempty"`
	Value   string `xml:",innerxml"`
}
type UniqueID struct {
	Default bool
	Value   string
}

func main() {
	//get()
	//write()
	scan()
}
func get() {
	client, err := tmdb.Init("e2fde5b39a4a7ae62fffaa548ea1b066")
	if err != nil {
		return
	}
	client.SetClientAutoRetry()
	client.SetAlternateBaseURL()
	//proxyUrl, err := url.Parse("http://localhost:7890")
	if err != nil {
		log.Println(err)
		return
	}
	customClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
			//Proxy:           request.ProxyURL(proxyUrl),
		},
	}
	client.SetClientConfig(customClient)
	urlOptions := map[string]string{
		"language":           "zh-CN",
		"append_to_response": "credits,images",
	}
	movie, err := client.GetTVDetails(225008, urlOptions)
	if err != nil {
		fmt.Println(err)
	}
	marshal, err := json.Marshal(&movie)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func write() {
	tv := TVShow{
		Title:         "The Big Bang Theory",
		OriginalTitle: "The Big Bang Theory",
		SortTitle:     "The Big Bang Theory",
		Ratings: Ratings{
			Rating: []Rating{
				Rating{
					Name:    "themoviedb",
					Max:     10,
					Default: true,
					Value:   154,
					Votes:   11,
				},
			},
		},
		Thumb: []Thumb{
			Thumb{
				Value:  "ASAD",
				Aspect: "1.5",
				Type:   "1",
				Season: 1,
			},
		},
	}
	marshal, err := xml.Marshal(&tv)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(xml.Header)
	fmt.Println(string(marshal))

}

func scan() {
	dir, err := os.ReadDir("../")
	if err != nil {
		return
	}
	for _, dirEntry := range dir {
		fmt.Println(dirEntry.Name())
	}
}
