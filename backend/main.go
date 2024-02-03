package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	data "github.com/KerenBermeo/CorreoQuery/extractAndProcessData.go"
	//i "github.com/KerenBermeo/CorreoQuery/indexZincsearch"
)

func main() {
	// Elimina el índice existente y crea uno nuevo
	// i.DeleteIndex()
	// i.CreateIndex()

	// Ruta del directorio raíz
	rootDirectory := "/Users/user/Desktop/EmailQuery/data"
	// rootDirectory := "/Users/user/Desktop/EmailQuery/data/blair-l/contacts"

	// Tamaño del lote (ajusta según sea necesario)
	batchSize := 500

	// Canal para almacenar archivos con un tamaño bufferizado
	filesChannel := make(chan []byte, batchSize)

	// Canal para señalizar que todos los archivos han sido procesados
	doneProcessing := make(chan bool)

	// Canal para pasar los datos procesados a la gorutina que maneja la operación masiva
	processedDataChannel := make(chan []byte, batchSize)

	// Utiliza WaitGroup para esperar que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Utiliza filepath.Walk para recorrer los archivos y directorios
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error al caminar por el directorio:", err)
				return err
			}

			// Verifica si la ruta del archivo termina con "_"
			if strings.HasSuffix(path, "_") {
				wg.Add(1)
				// Inicia una nueva goroutine llamando a walkFunction para procesar el archivo o directorio encontrado.
				go data.WalkFunction(path, info, err, &wg, filesChannel, processedDataChannel)
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error al caminar por el directorio principal:", err)
			close(filesChannel)
			wg.Wait()
			return
		}
	}()

	// Espera a que todas las goroutines finalicen
	wg.Wait()

	// Cierra el canal después de esperar que finalicen todas las gorutinas
	//close(filesChannel)

	// Todas las gorutinas de procesamiento han terminado, señaliza que la operación masiva puede comenzar
	doneProcessing <- true

	// Gorutina para manejar la operación masiva una vez que todos los archivos han sido procesados
	go func() {
		// Espera la señal de que todos los archivos han sido procesados
		<-doneProcessing

		// Inicializa un slice para almacenar los datos procesados
		//var bulkDataSlice []string

		// Consumir los datos procesados del nuevo canal
		for {
			// Intenta recibir un dato del canal, si está cerrado, rompe el bucle
			jsonData, ok := <-processedDataChannel
			if !ok {
				break
			}

			fmt.Printf("jsonData: %v, ok: %v\n", jsonData, ok)

			// // Convierte los bytes a string
			// StrJsonData := string(jsonData)

			// // Agrega el JSON al slice de datos procesados
			// bulkDataSlice = append(bulkDataSlice, StrJsonData)
		}

		// Construir el lote de datos utilizando bulkDataSlice
		// bulkData := strings.Join(bulkDataSlice, "\n")

		// // Envía el lote de datos a ZincSearch
		// data.SendToZincSearch(bulkData)
	}()

	// Espera a que todas las goroutines finalicen
	wg.Wait()

	// Canal de sincronización para cerrar processedDataChannel
	closeProcessedDataChannel := make(chan struct{})

	// Goroutine adicional para cerrar processedDataChannel
	go func() {
		defer close(closeProcessedDataChannel)
		// Espera a que todas las goroutines terminen antes de cerrar processedDataChannel
		wg.Wait()
		close(processedDataChannel)
	}()

	// Todas las gorutinas de procesamiento han terminado, señaliza que la operación masiva puede comenzar
	close(doneProcessing)

	// Espera a que la goroutine de cierre de canales termine
	<-closeProcessedDataChannel
}
