package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword), err
}

// ValidatePassword 验证密码
func ValidatePassword(password, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
