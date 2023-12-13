package model

import "gorm.io/gorm"

type VideoFiles struct {
	gorm.Model
	FullPath   string `json:"full_path"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Suffix     string `json:"suffix"`
	Recognized bool
}
