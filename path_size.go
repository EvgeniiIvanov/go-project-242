package goproject242

import (
	"os"
)

func GetSize(p string) (int64, error) {
	object, err := os.Lstat(p)
	if err != nil {
		return 0, err
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
		fileSize, err := GetSize(p + "/" + file.Name())
		if err != nil {
			// return error if we can't get size even for one file in dir
			return 0, err
		}
		size += fileSize
	}
	return size, nil

}
