package data

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// walkFunction recorre el directorio, filtra y procesa los archivos
func WalkFunction(path string, info os.FileInfo, err error, wg *sync.WaitGroup, filesChannel chan<- []byte, processedDataChannel chan<- []byte) {
	defer wg.Done()

	fmt.Println(filesChannel)
	//fmt.Println(processedDataChannel)

	// Manejo de errores al acceder al archivo o directorio
	if err != nil {
		fmt.Println("Error al acceder al archivo:", err, path)

		return
	}

	// Filtrar solo archivos regulares, ignorar directorios u otros tipos
	if !info.Mode().IsRegular() {
		return
	}

	// Limpiar la ruta del archivo
	cleanPath := filepath.Clean(path)

	// Leer el contenido del archivo
	content, err := os.ReadFile(cleanPath)

	if err != nil {
		// Manejo de errores al leer el archivo
		fmt.Printf("Error al leer el archivo %s: %v\n", cleanPath, err)
		return
	}

	// Procesar solo archivos, no directorios
	if !info.IsDir() {
		// Enviar el contenido del archivo al canal para procesar en lotes
		//filesChannel <- content

		// Llamada a ProcessBatch con la ruta del archivo actual
		jsonData, err := ProcessBatch(content, cleanPath)
		if err != nil {
			fmt.Println("Error al procesar el lote:", err)
			// Puedes manejar el error según tus necesidades
			return
		}

		processedDataChannel <- jsonData
	}
}
