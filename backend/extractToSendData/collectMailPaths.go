package data

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

// Devuelve una lista de paths de archivos encontrados.
func CollectMailsPaths(rootPath string) ([]string, error) {
	var slicePaths []string

	err := filepath.WalkDir(rootPath, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a %s: %v\n", path, err)
			return err // Devuelve el error encontrado
		}
		if !entry.IsDir() {
			slicePaths = append(slicePaths, path)
		}

		return nil // Continuar caminando
	})

	if err != nil {
		return nil, fmt.Errorf("error al caminar por el directorio: %v", err)
	}

	log.Println("Todos los paths han sido recopilados --> --> ")

	return slicePaths, nil
}
