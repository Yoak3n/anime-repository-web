package model

import "regexp"

type Rule struct {
	ID                uint
	VID               string
	Name              string
	Provider          string
	FileExtractReg    *regexp.Regexp
	Season            int
	Language          string
	EpisodeExtractReg *regexp.Regexp
	EpisodePosition   int
	EpisodeOffset     int
	Hidden            bool
}
