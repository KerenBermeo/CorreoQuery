package main

import (
	"fmt"
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"

	c "github.com/KerenBermeo/CorreoQuery/controllers"
	data "github.com/KerenBermeo/CorreoQuery/extractToSendData"
	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
	"github.com/KerenBermeo/CorreoQuery/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	getenv.ReadEnv()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	startTime := time.Now()

	passed, err := c.CheckIndexExists()
	if err != nil {
		log.Printf("error checking existence of index: %v", err)
	}

	rootDirectory := getenv.GetRootDirectory()
	if rootDirectory == "" {
		log.Println("Empty environment variable")
	}

	if passed {

		paths := data.GetAllFilePaths(rootDirectory)
		fmt.Println(len(paths))
		chunks := data.SplitIntoChunks(paths)

		var processFile model.FileProcessorFunc = data.ProcessFile
		resultChan := make(chan model.Email)
		var emails []model.Email

		data.OpenFiles(chunks, processFile, resultChan)

		for email := range resultChan {
			if email.Date != "" || email.From != "" || email.Subject != "" || email.Content != "" {
				emails = append(emails, email)
			}
		}

		data.CreateJSON(emails)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("total execution time: %s\n", duration)

	//Iniciar el servidor HTTP
	r := chi.NewRouter()
	router.ConfigureRouter(r)

	log.Fatal(http.ListenAndServe(":8081", r))
}
