package data

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CollectMailsPaths(rootPath string) []string {
	log.Print("Starting collection of mails paths")
	mailsPaths := []string{}
	err := filepath.Walk(rootPath,
		func(path string, fileInfo os.FileInfo, err error) error {
			// Use only files with no extension
			if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == "" {
				mailsPaths = append(mailsPaths, path)
			}
			return nil
		})
	if err != nil && err != io.EOF {
		log.Printf("Error while collecting paths: %s", err)
	}

	log.Printf("Mails path collection finalized. Total collected: %d", len(mailsPaths))
	return mailsPaths

}
