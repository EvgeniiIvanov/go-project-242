package goproject242

import (
	"fmt"
	"os"
)

func GetSize(p string, all bool) (int64, error) {
	object, err := os.Lstat(p)
	if err != nil {
		return 0, err
	}
	if !all {
		// skip hidden files started with .
		if object.Name()[0] == '.' {
			return 0, nil
		}
	}
	if !object.IsDir() {
		// calculate size of object if it is not a dir
		// Link is also file for us here
		return object.Size(), nil
	}
	// calculate size of files in dir
	files, err := os.ReadDir(p)
	if err != nil {
		return 0, err
	}
	//var size int64
	size := int64(0)
	for _, file := range files {
		// skip if this is dir
		if file.IsDir() {
			continue
		}
		// calculate size of file and add to total size
		fileSize, err := GetSize(p+"/"+file.Name(), all) // What if we use Windows?
		if err != nil {
			// return error if we can't get size even for one file in dir
			return 0, err
		}
		size += fileSize
	}
	return size, nil

}

func FormatSize(size int64, human bool) string {
	// If human is false, return size as is
	if !human {
		return fmt.Sprintf("%vB", size)
	}
	// Else
	// Create array of prefixes
	prefixes := []string{"", "K", "M", "G", "T", "P", "E"}
	// Create array of sizes
	sizes := []int64{
		1,
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
	// If size is less than 1 (e.g., empty file), return size as is
	return fmt.Sprintf("%vB", size)
}
