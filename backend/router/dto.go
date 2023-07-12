package router

import "github/Yoak3n/anime-repository-web/backend/model"

type Rule struct {
	ID                int    `json:"id"`
	VID               string `json:"vid"`
	Provider          string `json:"provider"`
	FileExtractReg    string `json:"file_extract_reg"`
	Season            int    `json:"season"`
	Language          string `json:"language"`
	EpisodeExtractReg string `json:"episode_extract_reg"`
	EpisodePosition   int    `json:"episode_position"`
	EpisodeOffset     int    `json:"episode_offset"`
}

func ExposeRule(rule *model.Rule) Rule {
	return Rule{
		ID:                rule.ID,
		VID:               rule.VID,
		Provider:          rule.Provider,
		Season:            rule.Season,
		Language:          rule.Language,
		EpisodePosition:   rule.EpisodePosition,
		EpisodeOffset:     rule.EpisodeOffset,
		FileExtractReg:    rule.FileExtractReg.String(),
		EpisodeExtractReg: rule.EpisodeExtractReg.String(),
	}
}
