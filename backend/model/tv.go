package model

import (
	"encoding/xml"
	"strings"
)

type TVShow struct {
	XMLName       xml.Name `xml:"tvshow"`
	Title         string   `xml:"title"`
	OriginalTitle string   `xml:"originaltitle"`
	//SortTitle     string
	Ratings   Ratings `xml:"ratings"`
	Premiered string  `xml:"premiered"`
	//UserRating string
	//Top250     int
	//Season     int
	//Episode    int
	Plot   string   `xml:"plot"`
	Studio []string `xml:"studio"`
	Genre  []string `xml:"genre"`
	//Tagline    string
	Thumb  []Thumb `xml:"thumb"`
	Fanart []Thumb `xml:"fanart"`
	//Mpaa       string
	//PlayCount  int
	//LastPlayed string
	UniqueID   UniqueID     `xml:"uniqueid"`
	Actor      []Actor      `xml:"actor"`
	NameSeason []NameSeason `xml:"nameseason"`
	DateAdded  string       `xml:"dateadded"`
}

type TVCollection struct {
	ClearLogo string `json:"clearlogo.png"`
	Poster    string `json:"poster.jpg"`
	Fanart    string `json:"fanart.jpg"`
	Seasons   map[string]string
}

type TVItem struct {
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

func NewTVItem(fullPath string) *TVItem {
	full := strings.Split(fullPath, "/")
	l := len(full)
	tvItem := &TVItem{}
	tvItem.FullPath = fullPath
	tvItem.Name = full[l-1]
	//tvItem.Path = strings.Join(full[:l-1], "/")
	tvItem.Path = strings.Replace(fullPath, tvItem.Name, "", 1)
	return tvItem
}

type TVEpisode struct {
	XMLName        xml.Name `xml:"episodedetails"`
	Title          string   `xml:"title"`
	OriginalTitle  string   `xml:"originaltitle"`
	ShowTitle      string   `xml:"showtitle"`
	Season         int64    `xml:"season"`
	Episode        int64    `xml:"episode"`
	DisplaySeason  int64    `xml:"displayseason"`
	Plot           string   `xml:"plot"`
	DisplayEpisode int64    `xml:"displayepisode"`
	UniqueID       UniqueID `xml:"uniqueid"`
	Aired          string   `xml:"aired"`
	Ratings        Ratings  `xml:"ratings"`
	Thumb          Thumb    `xml:"thumb"`
	Actor          []Actor  `xml:"actor"`
	DateAdded      string   `xml:"dateadded"`
}
