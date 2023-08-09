package main

import (
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/backend/database"
	"github/Yoak3n/anime-repository-web/backend/router"
)

func main() {
	defer database.GetConn().Close()
	go backthread.ScanList()
	router.RunSever()
}
