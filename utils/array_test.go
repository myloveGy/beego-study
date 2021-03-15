package utils

import (
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

func TestInArray(t *testing.T) {
	assert.Equal(t, true, InArray([]string{"username", "age"}, "username"))
	assert.Equal(t, false, InArray([]string{"username", "age"}, "username-1"))
}
