package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetPathSize for file
func TestGetPathSizeFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/hello.txt", false, false, false)

	a.NoError(err)
	a.Equal("14B", size)
}

// Test GetPathSize for dir
func TestGetPathSizeDir(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata", false, false, false)

	a.NoError(err)
	a.Equal("34B", size)
}

// Test GetPathSize for empty file
func TestGetPathSizeEmptyFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/empty.txt", false, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for non-existent file
func TestGetPathSizeNonExistent(t *testing.T) {
	a := assert.New(t)

	_, err := GetPathSize("./testdata/non-existent.txt", false, false, false)

	a.Error(err)
}

// Test GetPathSize for subdir
func TestGetPathSizeSubdir(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/subdir", false, false, false)

	a.NoError(err)
	a.Equal("21B", size)
}

// Test GetPathSize for empty subdir
func TestGetPathSizeEmptySubdir(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/subdir/empty", false, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for link
func TestGetPathSizeLink(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/goodbye.link", false, false, false)

	a.NoError(err)
	a.Equal("20B", size)
}

// Test GetPathSize for hidden file
func TestGetPathSizeHiddenFile(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden.txt", false, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for hidden file with all flag
func TestGetPathSizeHiddenFileAll(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden.txt", false, false, true)

	a.NoError(err)
	a.Equal("62B", size)
}

// Test GetPathSize for hidden dir
func TestGetPathSizeHiddenDir(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden", false, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for hidden dir with all flag
func TestGetPathSizeHiddenDirAll(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden", false, false, true)

	a.NoError(err)
	a.Equal("115B", size)
}

// Test GetPathSize for dir recursively
func TestGetPathSizeDirRecursive(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata", true, false, false)

	a.NoError(err)
	a.Equal("55B", size)
}

// Test GetPathSize for dir recursively with all flag
func TestGetPathSizeDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata", true, false, true)

	a.NoError(err)
	a.Equal("232B", size)
}

// Test GetPathSize for empty dir recursively
func TestGetPathSizeEmptyDirRecursive(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/empty", true, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for hidden dir recursively
func TestGetPathSizeHiddenDirRecursive(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden", true, false, false)

	a.NoError(err)
	a.Equal("0B", size)
}

// Test GetPathSize for hidden dir recursively with all flag
func TestGetPathSizeHiddenDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	size, err := GetPathSize("./testdata/.hidden", true, false, true)

	a.NoError(err)
	a.Equal("115B", size)
}

// Test FormatSize for bytes
func TestFormatSizeBytes(t *testing.T) {
	a := assert.New(t)

	size := int64(42)
	format := FormatSize(size)

	a.Equal("42B", format)
}

// Test FormatSize for KB
func TestFormatSizeKb(t *testing.T) {
	a := assert.New(t)

	size := int64(2048)
	format := FormatSize(size)

	a.Equal("2KB", format)
}

// Test FormatSize for MB
func TestFormatSizeMb(t *testing.T) {
	a := assert.New(t)

	size := int64(2097152)
	format := FormatSize(size)

	a.Equal("2MB", format)
}
