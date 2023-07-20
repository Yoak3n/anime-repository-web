package model

import "regexp"

type Rule struct {
	ID                int
	VID               string
	Provider          string
	FileExtractReg    *regexp.Regexp
	Season            int
	Language          string
	EpisodeExtractReg *regexp.Regexp
	EpisodePosition   int
	EpisodeOffset     int
	Hidden            bool
}
