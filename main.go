package main

import (
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/backend/router"
)

func main() {
	go backthread.ScanList()
	router.RunSever()
}
