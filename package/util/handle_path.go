package util

import (
	"fmt"
	"io"
	"os"
	"strings"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func MakeImagePath(imagePath string) string {
	return tmdb.GetImageURL(imagePath, tmdb.Original)
}

func CreateDirNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		e := os.MkdirAll(dir, 0777)
		if e != nil {
			println("Error creating directory: " + e.Error())
			return
		}
	}
}
func ReplacePressDown(original string) string {
	original, _ = strings.CutSuffix(original, "\\")
	original, _ = strings.CutSuffix(original, "/")
	return strings.Replace(original, "\\", "/", -1)
}

func MoveVideoFile(filePath string, newFilePath string) error {
	new_file, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("创建目标文件失败:", err)
		return err
	}
	defer new_file.Close()

	srcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开源文件失败:", err)
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(new_file, srcFile)
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	err = os.Symlink(newFilePath, filePath)
	if err != nil {
		return err
	}
	return nil
}
