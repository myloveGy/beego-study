package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	password, err := GeneratePassword("123456")
	assert.NoError(t, err)
	assert.Equal(t, true, password != "")
}

func TestValidatePassword(t *testing.T) {
	assert.Equal(t, true, ValidatePassword("123456", "$2a$10$ifs/VyyK6HzQUyyGoTu7D.vb036kFdIOBO.s3JevL/M03nC.1IS3W"))
}
