package response

import (
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
)

type Rule struct {
	ID                uint   `json:"id"`
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

type Config struct {
	ApiKey    string `json:"key"`
	UseProxy  bool   `json:"use_proxy"`
	Proxy     string `json:"proxy_url"`
	MoviePath string `json:"movie_path"`
	TVPath    string `json:"tv_path"`
	RawPath   string `json:"raw_path"`
	Delay     int64  `json:"delay"`
}

func ExposeConfig(conf *config.Configuration) Config {
	return Config{
		ApiKey:    conf.ApiKey,
		UseProxy:  conf.UseProxy,
		Proxy:     conf.Proxy.String(),
		MoviePath: conf.MoviePath,
		TVPath:    conf.TVPath,
		RawPath:   conf.RawPath,
		Delay:     conf.Delay,
	}
}
