package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// 创建一个文件夹
func MakeDir(folder string) error {
	folder, err := NormalizePath(folder)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(folder, 0750); err != nil {
		return err
	}
	return nil
}

//将文件中的~/xx 替换为绝对路径
func NormalizePath(path string) (string, error) {
	var err error
	if strings.HasPrefix(path, "~") {
		path, err = homedir.Expand(path)
		if err != nil {
			return path, err
		}
	}
	return path, nil
}

// FolderExists check if file is exist or not
func FolderExists(foldername string) bool {
	foldername, err := NormalizePath(foldername)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		fmt.Println(err)
		return false
	}
	return true
}

// EncodeBase64 returns a encoded string to base64
func EncodeBase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func InSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
