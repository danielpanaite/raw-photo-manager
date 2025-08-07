package backend

import (
	"fmt"
	"os"
	"strings"
)

var (
	pathPrefix string
)

func ListFiles(dir string) []string {
	pathPrefix = dir
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	files := []string{}
	for _, entry := range entries {
		parts := strings.Split(entry.Name(), ".")
		name := parts[0]
		ext := parts[1]
		if ext == "JPG" || ext == "jpg" || ext == "jpeg" {
			files = append(files, name)
		}
	}

	return files
}

func DeleteFiles(files []string) error {
	for _, file := range files {
		err := os.Remove(fmt.Sprintf("%s//%s.JPG", pathPrefix, file))
		if err != nil {
			return err
		}
		err = os.Remove(fmt.Sprintf("%s//%s.ARW", pathPrefix, file))
		if err != nil {
			return err
		}
	}
	return nil
}
