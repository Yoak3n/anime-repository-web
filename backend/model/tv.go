package model

import "encoding/xml"

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
