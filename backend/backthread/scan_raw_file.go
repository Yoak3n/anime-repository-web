package backthread

import (
	"errors"
	"fmt"
	"github/Yoak3n/anime-repository-web/backend/model"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/logger"
	"os"
	"strings"
	"time"
)

var delay int64

func ScanList() {
	for {
		time.Sleep(time.Second * time.Duration(delay))
		delay = config.Conf.Delay
		err := Scan()
		if err != nil {
			logger.ERROR.Println(err)
			continue
		}
		err = WorkOnMatchedCache()
		if err != nil {
			logger.ERROR.Println(err)
		}
	}
}

func Scan() error {
	if config.Conf.RawPath == "" {
		return errors.New("指定文件目录不能为空")
	}
	_, err := os.Stat(config.Conf.RawPath)
	if err != nil {
		return errors.New("指定文件目录不存在")
	}

	files, err := scanLoop(config.Conf.RawPath, "")
	videoFiles := make([]string, 0)
	for _, file := range files {
		if checkVideo(file) {
			logger.INFO.Println("找到文件", file)
			videoFiles = append(videoFiles, file)
		}
	}
	for _, file := range videoFiles {
		Dispose(file)
	}
	return nil
}

func GetFiles() (files []model.TVItem) {
	for _, tvItem := range cache.TvFiles {
		if !cache.Recognized[tvItem.FullPath] {
			files = append(files, *tvItem)
		}
	}
	return files
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
