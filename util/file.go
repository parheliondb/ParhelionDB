package util

import (
	"fmt"
	"os"
)

func CreateDirectoryIfNotExists(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return os.Mkdir(path, 0660)
	}

	if !fi.Mode().IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return nil
}
