package goproject242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSizeFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/hello.txt")

	a.NoError(err)
	a.Equal(int64(14), size)
}

func TestGetSizeDir(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata")

	a.NoError(err)
	a.Equal(int64(34), size)
}

func TestGetSizeEmptyFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/empty.txt")

	a.NoError(err)
	a.Equal(int64(0), size)
}

func TestGetSizeNonExistent(t *testing.T) {
	a := assert.New(t)

	_, err := GetSize("./testdata/non-existent.txt")

	a.Error(err)
}

func TestGetSizeSubdir(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/subdir")

	a.NoError(err)
	a.Equal(int64(21), size)
}
