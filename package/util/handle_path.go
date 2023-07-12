package util

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"os"
	"strings"
)

func MakeImagePath(imagePath string) string {
	return tmdb.GetImageURL(tmdb.Original, imagePath)
}

func CreateDirNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		e := os.Mkdir(dir, 0755)
		if e != nil {
			return
		}
	}
}
func ReplacePressDown(original string) string {
	original, _ = strings.CutSuffix(original, "\\")
	original, _ = strings.CutSuffix(original, "/")
	return strings.Replace(original, "\\", "/", -1)
}
