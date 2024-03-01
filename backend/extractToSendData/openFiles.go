package data

import (
	"log"
	"os"

	"sync"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func worker(inputChan <-chan *os.File, processFile model.FileProcessorFunc) model.Email {
	var email model.Email
	for file := range inputChan {
		email = processFile(file)
	}
	return email
}

func OpenFiles(chunks [][]string, processFile model.FileProcessorFunc, resultChan chan model.Email) {
	var wgWorkers sync.WaitGroup
	var wgFiles sync.WaitGroup
	var wgSenders sync.WaitGroup
	outputChan := make(chan *os.File, 3)
	numWorkers := 10

	for i := 0; i < numWorkers; i++ {
		wgWorkers.Add(1)
		wgSenders.Add(1)
		go func() {
			defer wgWorkers.Done()
			defer wgSenders.Done()
			email := worker(outputChan, processFile)
			//fmt.Println(email)
			resultChan <- email
		}()

	}

	sema := make(chan struct{}, 5)
	var filesToClose []*os.File

	for _, chunk := range chunks {
		for _, rootfile := range chunk {

			sema <- struct{}{}
			wgFiles.Add(1)
			go func(filename string) {
				defer func() { <-sema }()

				file, err := os.Open(filename)
				if err != nil {
					log.Println("error al abrir el archivo", err)
					return
				}

				fileInfo, err := file.Stat()
				if err != nil {
					log.Println("error al obtener información del archivo", err)
					file.Close()
					return
				}

				if fileInfo.Size() > 0 {
					outputChan <- file // <------- se envian los archivos abiertos al canal

					filesToClose = append(filesToClose, file)
				} else {
					file.Close()
				}

				wgFiles.Done()
			}(rootfile)
		}
	}

	// Esperar a que todas las gorutinas de los trabajadores terminen
	go func() {
		wgFiles.Wait()
		close(outputChan)

		batchSize := 5

		// Cerrar archivos en lotes
		for len(filesToClose) > 0 {
			batch := filesToClose
			if len(filesToClose) > batchSize {
				batch = filesToClose[:batchSize]
				filesToClose = filesToClose[batchSize:]
			} else {
				filesToClose = nil
			}
			for _, file := range batch {
				file.Close()
			}
		}

	}()

	go func() {
		wgWorkers.Wait()
		close(resultChan)
	}()

}
