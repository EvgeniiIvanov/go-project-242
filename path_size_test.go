package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetPathSize for file
func TestGetPathSizeFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/hello.txt", o)

	a.NoError(err)
	a.Equal(int64(14), size)
}

// Test GetPathSize for dir
func TestGetPathSizeDir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(34), size)
}

// Test GetPathSize for empty file
func TestGetPathSizeEmptyFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/empty.txt", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for non-existent file
func TestGetPathSizeNonExistent(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	_, err := GetPathSize("./testdata/non-existent.txt", o)

	a.Error(err)
}

// Test GetPathSize for subdir
func TestGetPathSizeSubdir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/subdir", o)

	a.NoError(err)
	a.Equal(int64(21), size)
}

// Test GetPathSize for empty subdir
func TestGetPathSizeEmptySubdir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/subdir/empty", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for link
func TestGetPathSizeLink(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/goodbye.link", o)

	a.NoError(err)
	a.Equal(int64(20), size)
}

// Test GetPathSize for hidden file
func TestGetPathSizeHiddenFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/.hidden.txt", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for hidden file with all flag
func TestGetPathSizeHiddenFileAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: true}

	size, err := GetPathSize("./testdata/.hidden.txt", o)

	a.NoError(err)
	a.Equal(int64(62), size)
}

// Test GetPathSize for hidden dir
func TestGetPathSizeHiddenDir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetPathSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for hidden dir with all flag
func TestGetPathSizeHiddenDirAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: true}

	size, err := GetPathSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(115), size)
}

// Test GetPathSize for dir recursively
func TestGetPathSizeDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetPathSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(55), size)
}

// Test GetPathSize for dir recursively with all flag
func TestGetPathSizeDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: true}

	size, err := GetPathSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(232), size)
}

// Test GetPathSize for empty dir recursively
func TestGetPathSizeEmptyDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetPathSize("./testdata/empty", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for hidden dir recursively
func TestGetPathSizeHiddenDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetPathSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetPathSize for hidden dir recursively with all flag
func TestGetPathSizeHiddenDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: true}

	size, err := GetPathSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(115), size)
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
