package backthread

import (
	"errors"
	"fmt"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"log"
	"os"
	"strings"
	"time"
)

var delay int64

func ScanList() {
	for {
		delay = config.Conf.Delay
		videoFiles, err := Scan()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * time.Duration(delay))
			continue
		}
		for _, file := range videoFiles {
			Dispose(file)
		}
		time.Sleep(time.Second * time.Duration(delay))
	}
}

func Scan() ([]string, error) {
	_, err := os.Stat(config.Conf.RawPath)
	if err != nil {
		return nil, errors.New("指定文件目录不存在")
	}

	files, err := scanLoop(config.Conf.RawPath, "")
	videoFiles := make([]string, 0)
	for _, file := range files {
		if checkVideo(file) {
			logger.INFO.Println("找到文件", file)
			videoFiles = append(videoFiles, file)
		}
	}
	return videoFiles, nil
}

func scanLoop(path string, name string) (videoFiles []string, err error) {
	var subs []os.DirEntry
	var currentPath string
	if name != "" {
		currentPath = fmt.Sprintf(`%s/%s`, path, name)
		logger.INFO.Println("下级目录", currentPath)
	} else {
		currentPath = path
		logger.INFO.Println("主目录", currentPath)
	}

	subs, err = os.ReadDir(currentPath)
	if err != nil {
		return nil, err
	}
	for _, sub := range subs {
		if sub.IsDir() {
			var loop []string
			loop, err = scanLoop(currentPath, sub.Name())
			if err != nil {
				return nil, err
			}
			videoFiles = append(videoFiles, loop...)
		} else {
			videoFiles = append(videoFiles, currentPath+"/"+sub.Name())
		}
	}
	return videoFiles, nil
}

func checkVideo(file string) bool {
	videoType := []string{".mp4", ".mkv", ".avi"}
	for _, t := range videoType {
		if strings.HasSuffix(file, t) {
			return true
		}
	}
	return false
}
