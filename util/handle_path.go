package util

import tmdb "github.com/cyruzin/golang-tmdb"

func makeImagePath(imagePath string) string {
	return tmdb.GetImageURL(tmdb.Original, imagePath)
}
