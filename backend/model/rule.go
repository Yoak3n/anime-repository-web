package model

import (
	"gorm.io/gorm"
	"regexp"
)

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

type RuleModel struct {
	gorm.Model
	VID               string `gorm:"column:vid"`
	Name              string `gorm:"column:name"`
	Provider          string `gorm:"column:provider"`
	FileExtractReg    string `gorm:"column:file_extract_reg"`
	Season            int    `gorm:"column:season"`
	Language          string `gorm:"column:language"`
	EpisodeExtractReg string `gorm:"column:episode_extract_reg"`
	EpisodePosition   int    `gorm:"column:episode_position"`
	EpisodeOffset     int    `gorm:"column:episode_offset"`
}
