package main

import (
	"fmt"
	"sync"
	"time"

	_ "net/http/pprof"

	data "github.com/KerenBermeo/CorreoQuery/extractAndProcessData.go"
	index "github.com/KerenBermeo/CorreoQuery/indexZincsearch"
)

func main() {

	//Borrar y crear indice
	index.DeleteIndex()
	index.CreateIndex()

	// Utiliza WaitGroup para esperar que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Registra el tiempo actual antes de comenzar la ejecución del programa
	startTime := time.Now()

	// Ruta del directorio raíz
	// rootDirectory := "/Users/user/Desktop/EmailQuery/data/allen-p"
	rootDirectory := "/Users/user/Desktop/archivos"

	// Tamaño del lote (ajusta según sea necesario)
	batchSize := 100

	// Path de dirreciones de los archivos
	mailsPaths := data.CollectMailsPaths(rootDirectory)

	chunks := data.ChunkEmails(mailsPaths)

	wg.Add(len(chunks))
	for _, chunk := range chunks {
		go data.ProcessBatch(chunk, batchSize, &wg)
	}
	wg.Wait()

	// Registra el tiempo después de que todas las goroutines han terminado
	endTime := time.Now()

	// Calcula y muestra la duración total de la ejecución
	duration := endTime.Sub(startTime)
	fmt.Printf("Tiempo total de ejecución: %s\n", duration)

}
