package data

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// CollectMailsPaths recopila los paths de archivos sin extensión en el directorio dado.
// Utiliza concurrencia para procesar los archivos de manera eficiente.
// Devuelve una lista de paths de archivos encontrados.
func CollectMailsPaths(rootPath string) ([]string, error) {
	var mailsPaths sync.Map // Usamos sync.Map en lugar de []string
	var wg sync.WaitGroup
	var semaphore = make(chan struct{}, 20)
	var filesCh = make(chan string)

	// Función que representa a los trabajadores
	worker := func() {
		defer wg.Done()
		for path := range filesCh {
			mailsPaths.Store(path, true) // Almacenamos el path en sync.Map
			<-semaphore                  // Libera un espacio en el semáforo
		}
	}

	// Iniciar trabajadores
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker()
	}

	// Función para caminar por los archivos y enviarlos al canal filesCh
	walkFunc := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err // Manejar errores específicos aquí si es necesario
		}
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == "" {
			semaphore <- struct{}{} // Ocupa un espacio en el semáforo
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

	// Convertir el sync.Map a un slice de strings
	var result []string
	mailsPaths.Range(func(key, value interface{}) bool {
		result = append(result, key.(string))
		return true
	})

	return result, nil
}
