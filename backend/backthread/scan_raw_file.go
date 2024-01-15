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
	if err != nil {
		return err
	}
	videoFiles := make([]string, 0)
	for _, file := range files {
		if checkVideo(file) {
			videoFiles = append(videoFiles, file)
		}
	}
	for _, file := range videoFiles {
		Dispose(file)
	}
	err = WorkOnMatchedCache()
	if err != nil {
		logger.ERROR.Println(err)
	}
	MatchRules()
	return nil
}

func GetFiles() (files []model.TVItem) {
	err := Scan()
	if err != nil {
		logger.ERROR.Println(err)
		return files
	}
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
	} else {
		currentPath = path
	}

	subs, err = os.ReadDir(currentPath)
	if err != nil {
		return nil, err
	}
	for _, sub := range subs {
		fileInfo, err := os.Lstat(currentPath + "/" + sub.Name())
		if err != nil {
			logger.DEBUG.Println("无法获取文件信息")
			continue
		}
		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			logger.DEBUG.Println("文件是软链接，跳过")
			continue
		}
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
