package util

import (
	"fmt"
	"os"
)

func CreateDirectoryIfNotExists(path string) error {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return os.Mkdir(path, 0660)
	} else if !fi.Mode().IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return nil
}
