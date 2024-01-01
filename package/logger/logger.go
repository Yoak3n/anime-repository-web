package logger

import (
	"github/Yoak3n/anime-repository-web/package/util"
	"io"
	"log"
	"os"
)

var (
	INFO  *log.Logger
	DEBUG *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
)

func init() {
	util.CreateDirNotExists("data/logs")
	file, err := os.OpenFile("data/logs/errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Can't open file error.log:", err)
	}
	defer file.Close()
	INFO = log.New(log.Writer(), "[INFO] ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
	DEBUG = log.New(log.Writer(), "[DEBUG] ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
	WARN = log.New(log.Writer(), "[WARN] ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
	ERROR = log.New(io.MultiWriter(file, log.Writer()), "[ERROR] ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
}
