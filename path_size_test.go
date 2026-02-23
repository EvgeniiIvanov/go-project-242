package goproject242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetSize for file
func TestGetSizeFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/hello.txt")

	a.NoError(err)
	a.Equal(int64(14), size)
}

// Test GetSize for dir
func TestGetSizeDir(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata")

	a.NoError(err)
	a.Equal(int64(34), size)
}

// Test GetSize for empty file
func TestGetSizeEmptyFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/empty.txt")

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for non-existent file
func TestGetSizeNonExistent(t *testing.T) {
	a := assert.New(t)

	_, err := GetSize("./testdata/non-existent.txt")

	a.Error(err)
}

// Test GetSize for subdir
func TestGetSizeSubdir(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/subdir")

	a.NoError(err)
	a.Equal(int64(21), size)
}

// Test GetSize for empty subdir
func TestGetSizeEmptySubdir(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/subdir/empty")

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for link
func TestGetSizeLink(t *testing.T) {
	a := assert.New(t)

	size, err := GetSize("./testdata/goodbye.link")

	a.NoError(err)
	a.Equal(int64(20), size)
}

// Test FormatSize for bytes
func TestFormatSizeBytes(t *testing.T) {
	a := assert.New(t)

	size := int64(42)
	format := FormatSize(size, true)

	a.Equal("42.0B", format)
}

// Test FormatSize for KB
func TestFormatSizeKb(t *testing.T) {
	a := assert.New(t)

	size := int64(2048)
	format := FormatSize(size, true)

	a.Equal("2.0KB", format)
}

// Test FormatSize for MB
func TestFormatSizeMb(t *testing.T) {
	a := assert.New(t)

	size := int64(2097152)
	format := FormatSize(size, true)

	a.Equal("2.0MB", format)
}

// Test FormatSize for false flag
func TestFormatSizeFalse(t *testing.T) {
	a := assert.New(t)

	size := int64(2048)
	format := FormatSize(size, false)

	a.Equal("2048B", format)
}
