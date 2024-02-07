package data

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// CollectMailsPaths recopila los paths de archivos de correos electrónicos
// en el directorio raíz especificado por rootPath.
func CollectMailsPaths(rootPath string) []string {
	// mailsPaths almacena los paths de los archivos de correos electrónicos.
	mailsPaths := []string{}

	// Recorre el directorio raíz para encontrar archivos de correos electrónicos.
	err := filepath.Walk(rootPath,
		func(path string, fileInfo os.FileInfo, err error) error {
			// Solo se incluyen archivos sin extensión.
			if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == "" {
				mailsPaths = append(mailsPaths, path)
			}
			return nil
		})

	// Maneja errores durante la recopilación de paths de archivos.
	if err != nil && err != io.EOF {
		log.Printf("Error al recopilar paths: %s", err)
	}

	return mailsPaths
}
