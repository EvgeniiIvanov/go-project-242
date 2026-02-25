package goproject242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetSize for file
func TestGetSizeFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/hello.txt", o)

	a.NoError(err)
	a.Equal(int64(14), size)
}

// Test GetSize for dir
func TestGetSizeDir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(34), size)
}

// Test GetSize for empty file
func TestGetSizeEmptyFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/empty.txt", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for non-existent file
func TestGetSizeNonExistent(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	_, err := GetSize("./testdata/non-existent.txt", o)

	a.Error(err)
}

// Test GetSize for subdir
func TestGetSizeSubdir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/subdir", o)

	a.NoError(err)
	a.Equal(int64(21), size)
}

// Test GetSize for empty subdir
func TestGetSizeEmptySubdir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/subdir/empty", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for link
func TestGetSizeLink(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/goodbye.link", o)

	a.NoError(err)
	a.Equal(int64(20), size)
}

// Test GetSize for hidden file
func TestGetSizeHiddenFile(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/.hidden.txt", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for hidden file with all flag
func TestGetSizeHiddenFileAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: true}

	size, err := GetSize("./testdata/.hidden.txt", o)

	a.NoError(err)
	a.Equal(int64(62), size)
}

// Test GetSize for hidden dir
func TestGetSizeHiddenDir(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: false}

	size, err := GetSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for hidden dir with all flag
func TestGetSizeHiddenDirAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: false, All: true}

	size, err := GetSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(115), size)
}

// Test GetSize for dir recursively
func TestGetSizeDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(55), size)
}

// Test GetSize for dir recursively with all flag
func TestGetSizeDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: true}

	size, err := GetSize("./testdata", o)

	a.NoError(err)
	a.Equal(int64(232), size)
}

// Test GetSize for empty dir recursively
func TestGetSizeEmptyDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetSize("./testdata/empty", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for hidden dir recursively
func TestGetSizeHiddenDirRecursive(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: false}

	size, err := GetSize("./testdata/.hidden", o)

	a.NoError(err)
	a.Equal(int64(0), size)
}

// Test GetSize for hidden dir recursively with all flag
func TestGetSizeHiddenDirRecursiveAll(t *testing.T) {
	a := assert.New(t)

	o := Options{Recursive: true, All: true}

	size, err := GetSize("./testdata/.hidden", o)

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
