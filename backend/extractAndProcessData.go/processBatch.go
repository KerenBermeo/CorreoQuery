package data

import (
	"log"
	"sync"
)

// ProcessBatch procesa un lote de rutas de archivos de correo electrónico, los parsea y los guarda en lotes en la base de datos.
// La función toma un slice de rutas de archivos de correo electrónico y un tamaño de lote como entrada.
func ProcessBatch(mailsPaths []string, bulkSize int, wg *sync.WaitGroup) {
	// Obtiene la cantidad total de rutas de archivos de correo electrónico en el lote.
	total := len(mailsPaths)

	//indice a agregar
	strIndexName := `{ "index" : { "_index" : "olympics" } }` + "\n"

	// Registra información sobre la preparación para procesar los archivos.
	log.Printf("Preparándose para procesar archivos. Tamaño del lote: %v. Total de registros: %v", bulkSize, total)

	// Inicializa un slice para almacenar los correos electrónicos parseados en el lote actual.
	// var bulk []string
	var bulk []string

	// Itera sobre cada ruta de archivo en el lote.
	for i, item := range mailsPaths {
		// Parsea el correo electrónico a partir de la ruta del archivo.
		parsedEmail, err := ParseEmailFromPath(item)

		// Si hay un error al parsear, continúa con la siguiente iteración.
		if err != nil {
			continue
		}

		parseJson, err := ParsedEmail(parsedEmail)
		if err != nil {
			continue
		}

		strIndexName = strIndexName + parseJson

		// Agrega el correo electrónico parseado al lote.
		bulk = append(bulk, strIndexName)

		// Si el tamaño actual del lote alcanza el tamaño especificado, lo guarda en la base de datos y reinicia el lote.
		if (i+1)%bulkSize == 0 {
			log.Printf("Subiendo lote %v / %v", i+1, total)
			SendToZincSearch(bulk)
			bulk = nil
		} else if bulk != nil && (i+1) == total {
			// Si es la última iteración y aún hay elementos en el lote, lo guarda en la base de datos.
			log.Printf("Subiendo último lote %v / %v", i+1, total)
			SendToZincSearch(bulk)
		}
	}
	// Marca la finalización de la goroutine actual utilizando WaitGroup.
	wg.Done()
}
