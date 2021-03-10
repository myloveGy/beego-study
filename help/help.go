package help

import (
	"os"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func MapArray(arr []interface{}, key, val string) map[string]interface{} {
	m := make(map[string]interface{})
	for i, iL := 0, len(arr); i < iL; i++ {
		tmp := reflect.ValueOf(arr[i])
		id := strconv.FormatInt(tmp.FieldByName(key).Int(), 10)
		m[id] = tmp.FieldByName(val).Interface()
	}

	return m
}

// IsFileExists 判断文件是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//  IsDirExists 判断目录是否存在
func IsDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return file.IsDir()
}

// 判断文件是否在一个数组中
func InArray(arr []string, val string) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}

	return false
}

// 获取文件大小的接口
type Sizer interface {
	Size() int64
}

// GeneratePassword 生成密码
func GeneratePassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword), err
}

// ValidatePassword 验证密码
func ValidatePassword(password, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
