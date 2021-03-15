package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExists(t *testing.T) {
	exists := IsFileExists("./array.go")
	assert.Equal(t, true, exists)
	assert.Equal(t, false, IsFileExists("./user.go"))
}

func TestIsDirExists(t *testing.T) {
	exists := IsDirExists("./")
	assert.Equal(t, true, exists)
	assert.Equal(t, false, IsDirExists("./index"))
}
