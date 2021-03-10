package help

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type student struct {
	Id   int64
	Name string
}

func TestMapArray(t *testing.T) {
	v := []interface{}{
		student{1, "username"},
		student{2, "jinxing.liu"},
	}

	m := MapArray(v, "Id", "Name")
	assert.Equal(t, map[string]interface{}{
		"1": "username",
		"2": "jinxing.liu",
	}, m)
}

func TestIsFileExists(t *testing.T) {
	exists := IsFileExists("./help.go")
	assert.Equal(t, true, exists)
	assert.Equal(t, false, IsFileExists("./user.go"))
}

func TestIsDirExists(t *testing.T) {
	exists := IsDirExists("./")
	assert.Equal(t, true, exists)
	assert.Equal(t, false, IsDirExists("./index"))
}

func TestInArray(t *testing.T) {
	assert.Equal(t, true, InArray([]string{"username", "age"}, "username"))
	assert.Equal(t, false, InArray([]string{"username", "age"}, "username-1"))
}

func TestGeneratePassword(t *testing.T) {
	password, err := GeneratePassword("123456")
	assert.NoError(t, err)
	assert.Equal(t, true, password != "")
	fmt.Println(password)
}

func TestValidatePassword(t *testing.T) {
	assert.Equal(t, true, ValidatePassword("123456", "$2a$10$ifs/VyyK6HzQUyyGoTu7D.vb036kFdIOBO.s3JevL/M03nC.1IS3W"))
}
