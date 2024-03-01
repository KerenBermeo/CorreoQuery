package data

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func GetAllFilePaths(root string) []string {
	var slicePaths []string

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return err
		}
		if !entry.IsDir() {
			slicePaths = append(slicePaths, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error when walking through the directory: %v\n", err)
	}

	return slicePaths

}
