package main

import (
	"log"
	"sync"
	"time"

	"net/http"
	_ "net/http/pprof"

	c "github.com/KerenBermeo/CorreoQuery/controllers"
	data "github.com/KerenBermeo/CorreoQuery/extractToSendData"
	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	getenv.ReadEnv()
	// Iniciar el servidor de perfilamiento de Go (pprof) en segundo plano
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Iniciar un temporizador para registrar el tiempo de inicio de la ejecución del programa
	startTime := time.Now()

	// Verificar si el índice existe
	passed, err := c.CheckIndexExists()
	if err != nil {
		log.Printf("error al verificar la existencia del índice: %v", err)
	}

	// Crear un WaitGroup para esperar a que todas las goroutines finalicen
	var wg sync.WaitGroup

	rootDirectory := getenv.GetRootDirectory()
	if rootDirectory == "" {
		log.Println("Variable de entorno vacia")
	}

	//tamaño del lote
	maxBatch := 1000

	if passed {

		// Obtener los paths de los archivos de correos electrónicos
		mailsPaths, err := data.CollectMailsPaths(rootDirectory)
		if err != nil {
			log.Printf("error al recopilar paths: %s", err)
		}

		//Dividir el slice de paths en chunks para su procesamiento
		chunks := data.ChunkEmails(mailsPaths)
		c.NumberFiles = len(mailsPaths)
		log.Println("Archivos: ", c.NumberFiles)

		data.Count = 0
		wg.Add(len(chunks))
		for _, chunk := range chunks {

			go data.ParseEmailFromPath(chunk, maxBatch, &wg)
		}
		wg.Wait()
	}

	// Registrar el tiempo después de que todas las goroutines han terminado
	endTime := time.Now()

	// Calcular y mostrar la duración total de la ejecución
	duration := endTime.Sub(startTime)

	log.Printf("tiempo total de ejecución: %s\n", duration)

	//Iniciar el servidor HTTP
	r := chi.NewRouter()
	router.ConfigureRouter(r)
	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, r))
}
