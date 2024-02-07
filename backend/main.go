package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"net/http"
	_ "net/http/pprof"

	data "github.com/KerenBermeo/CorreoQuery/extractToSendData"

	"github.com/KerenBermeo/CorreoQuery/router"
	"github.com/go-chi/chi/v5"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// startTime registra el tiempo de inicio de la ejecución del programa.
	startTime := time.Now()

	// Nombre del índice que deseas verificar
	indexName := "email"

	// verificar que el indice no exista, en caso de que exista eliminarlo.
	err := data.CheckIndexExists(indexName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// wg es un WaitGroup para esperar a que todas las goroutines finalicen.
	var wg sync.WaitGroup

	// rootDirectory es la ruta del directorio raíz que contiene los correos electrónicos.
	rootDirectory := "/Users/user/Desktop/EmailQuery/data"
	//rootDirectory := "/Users/user/Desktop/archivos"
	//rootDirectory := "/Users/user/Desktop/EmailQuery/data/arnold-j"

	// batchSize es el tamaño del lote de correos electrónicos a procesar en cada iteración.
	batchSize := 500

	// mailsPaths contiene los paths de los archivos de correos electrónicos.
	mailsPaths := data.CollectMailsPaths(rootDirectory)

	// chunks contiene lotes de paths de correos electrónicos para procesamiento concurrente.
	chunks := data.ChunkEmails(mailsPaths)

	// Agrega el número de chunks al WaitGroup.
	wg.Add(len(chunks))

	// Procesa cada chunk de correos electrónicos concurrentemente.
	for _, chunk := range chunks {
		go data.ProcessBatch(chunk, batchSize, &wg)
	}

	// Espera a que todas las goroutines finalicen.
	wg.Wait()

	// Registra el tiempo después de que todas las goroutines han terminado.
	endTime := time.Now()

	// Calcula y muestra la duración total de la ejecución.
	duration := endTime.Sub(startTime)
	fmt.Printf("Tiempo total de ejecución: %s\n", duration)

	r := chi.NewRouter()
	router.SetupRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
