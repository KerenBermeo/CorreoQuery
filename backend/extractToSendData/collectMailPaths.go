package data

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// Devuelve una lista de paths de archivos encontrados.
func CollectMailsPaths(rootPath string) ([]string, error) {
	var mailsPathsMu sync.Mutex
	mailsPaths := make(map[string]bool)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 32)
	filesCh := make(chan string) // Definir el canal filesCh

	log.Println("Inicio recoleccion de paths -->")

	// Función que representa a los trabajadores
	worker := func() {
		defer wg.Done()
		for path := range filesCh {
			mailsPathsMu.Lock()
			mailsPaths[path] = true
			mailsPathsMu.Unlock()
			<-semaphore
		}
	}

	// Iniciar trabajadores
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go worker()
	}

	// Función para caminar por los archivos y enviarlos al canal filesCh
	walkFunc := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == "" {
			semaphore <- struct{}{}
			filesCh <- path
		}
		return nil
	}

	// Realizar el recorrido del directorio con el walkFunc
	err := filepath.Walk(rootPath, walkFunc)

	// Cerrar el canal filesCh después de caminar el directorio
	close(filesCh)

	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("error al recopilar paths: %s", err)
	}

	wg.Wait()

	// Convertir el mapa a un slice de strings
	var result []string
	mailsPathsMu.Lock()
	for path := range mailsPaths {
		result = append(result, path)
	}
	mailsPathsMu.Unlock()

	log.Println("Todos los paths han sido recopilados --> --> ")

	return result, nil
}
