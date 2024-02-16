package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/KerenBermeo/CorreoQuery/controllers"
	data "github.com/KerenBermeo/CorreoQuery/extractToSendData"
	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
	// 	"github.com/KerenBermeo/CorreoQuery/router"
	// 	"github.com/go-chi/chi/v5"
)

func main() {
	// Iniciar el servidor de perfilamiento de Go (pprof) en segundo plano
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Iniciar un temporizador para registrar el tiempo de inicio de la ejecución del programa
	startTime := time.Now()

	// Obtener el índice de nombres desde algún lugar (aquí se asume que se obtiene del entorno)
	nameIndex := getenv.GetNameIndex()

	// Verificar si el índice existe en los controladores
	passed, err := controllers.CheckIndexExists(nameIndex)
	if err != nil {
		log.Printf("error al verificar la existencia del índice: %v", err)
	}

	// Crear un WaitGroup para esperar a que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Definir la ruta del directorio raíz que contiene los correos electrónicos
	//rootDirectory := "/Users/user/Desktop/EmailQuery/data/allen-p"
	rootDirectory := "/Users/user/Desktop/archivos_2"
	// rootDirectory := "/Users/user/Desktop/EmailQuery/data"

	// Si el índice existe, proceder con el procesamiento de correos electrónicos
	if passed {
		// Obtener los paths de los archivos de correos electrónicos
		mailsPaths, err := data.CollectMailsPaths(rootDirectory)
		if err != nil {
			log.Printf("error al recopilar paths: %s", err)
		}

		// Se determina el número de núcleos lógicos disponibles en el sistema.
		numCPU := runtime.NumCPU()

		// Se obtiene la cantidad total de rutas de correos electrónicos en la lista.
		numMails := len(mailsPaths)
		fmt.Println("numMails ------> ", numMails)
		//fmt.Println(numMails)
		fmt.Println("numCPU ------> ", numCPU)
		// Se calcula el tamaño de cada fragmento, considerando el número de núcleos lógicos.
		chunkSize := (numMails + numCPU - 1) / numCPU
		fmt.Println("chunkSize ------> ", chunkSize)
		//fmt.Println(chunkSize)

		var models []model.Email
		contador := 0
		for i := 0; i < len(mailsPaths); i += chunkSize {
			end := i + chunkSize
			contador++

			if end > len(mailsPaths) {
				end = len(mailsPaths)
			}
			wg.Add(1)
			go func(start, end int) {
				models, err = data.ParseEmailFromPath(contador, mailsPaths[start:end])

				if err != nil {
					log.Printf("error al parsear correos electrónicos: %v", err)
					return
				}

				fmt.Println("Inicio de transformacion a json ->")
				if (i+1)%chunkSize == 0 { // Upload bulk and start over
					log.Printf("Uploading bulk %v / %v", i+1, len(mailsPaths))
					data.ConcurrentParsedEmailJson(models)
					models = nil
				} else if models != nil && (i+1) == len(mailsPaths) { // Upload last bulk
					log.Printf("Uploading bulk %v / %v", i+1, len(mailsPaths))
					data.ConcurrentParsedEmailJson(models)
				}

			}(i, end)
		}

		wg.Wait()

		// Se encarga de transformar las estructuras a json

	}

	// Registrar el tiempo después de que todas las goroutines han terminado
	endTime := time.Now()

	// Calcular y mostrar la duración total de la ejecución
	duration := endTime.Sub(startTime)

	fmt.Printf("tiempo total de ejecución: %s\n", duration)

	// Iniciar el servidor HTTP con el enrutador configurado (si es necesario)
	// r := chi.NewRouter()
	// router.ConfigureRouter(r)
	// log.Fatal(http.ListenAndServe(":8080", r))
}
