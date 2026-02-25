package code

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	if human {
		return FormatSize(size), nil
	}
	return fmt.Sprintf("%dB", size), nil
}

func getSize(path string, all, recursive bool) (int64, error) {
	object, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !all {
		// skip hidden files started with ".", and files with name length <= 0
		if len(object.Name()) > 0 && object.Name()[0] == '.' {
			return 0, nil
		}
	}
	if !object.IsDir() {
		// calculate size of object if it is not a dir
		// Link is also file for us here
		return object.Size(), nil
	}
	// calculate size of files in dir
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	size := int64(0)
	for _, file := range files {
		if file.IsDir() && !recursive {
			continue
		}
		// calculate size of file/dir and add to total size
		fileSize, err := getSize(filepath.Join(path, file.Name()), all, recursive)
		if err != nil {
			// return error if we can't get size even for one file in dir
			return 0, err
		}
		size += fileSize
	}
	return size, nil
}

func FormatSize(size int64) string {
	// Create array of prefixes
	prefixes := []string{"K", "M", "G", "T", "P", "E"}
	// Create array of sizes
	sizes := []int64{
		1 << 10, // 1024
		1 << 20, // 1048576
		1 << 30, // 1073741824
		1 << 40, // 1099511627776
		1 << 50, // 1125899906842624
		1 << 60, // 1152921504606846976
	}

	// Compare size with values of sizes and add corresponding prefix
	for i := len(sizes) - 1; i >= 0; i-- {
		if size >= sizes[i] {
			return fmt.Sprintf("%.1f%sB", float64(size)/float64(sizes[i]), prefixes[i])
		}
	}
	// If size is less than 1024 bytes, return size as is
	return fmt.Sprintf("%vB", size)
}
