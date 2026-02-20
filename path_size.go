package goproject242

import (
	"os"
)

func GetSize(p string) (int64, error) {
	size, err := os.Lstat(p)
	if err != nil {
		return 0, err
	}
	return size.Size(), nil
}
