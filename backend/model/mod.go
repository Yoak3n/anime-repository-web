package model

type Ratings struct {
	Rating []Rating `xml:"rating"`
}
type Rating struct {
	Name    string  `xml:"name,attr"`
	Max     int     `xml:"max,attr"`
	Default bool    `xml:"default,attr"`
	Value   float32 `xml:"value"`
	Votes   int64   `xml:"votes"`
}
type Thumb struct {
	Aspect  string `xml:"aspect,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Season  int    `xml:"season,attr,omitempty"`
	Preview string `xml:"preview,omitempty"`
	Value   string `xml:",innerxml"`
}
type UniqueID struct {
	Default bool   `xml:"default,attr"`
	Type    string `xml:"type,attr"`
	Value   string `xml:",innerxml"`
}
type Actor struct {
	Name  string `xml:"name"`
	Role  string `xml:"role"`
	Order int    `xml:"order"`
	Thumb Thumb  `xml:"thumb"`
}
type NameSeason struct {
	Number int    `xml:"number,attr"`
	Value  string `xml:",innerxml"`
}
