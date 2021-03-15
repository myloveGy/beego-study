package utils

import (
	"os"
)

// 获取文件大小的接口
type Sizer interface {
	Size() int64
}

// IsFileExists 判断文件是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDirExists 判断目录是否存在
func IsDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return file.IsDir()
}
