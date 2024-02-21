package data

import (
	"log"
	"net/http"
	"strings"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

var Count int

func SendBatch(batch []string) {
	Count++

	log.Println("El lote numero: ", Count, " Tamaño del batch: ", len(batch))

	serverUrl := getenv.GetZincSearchServerURL()
	if serverUrl == "" {
		log.Println("Variable de entorno vacia")
	}

	url := serverUrl + "api/_bulk"

	// Crear la solicitud HTTP con los correos electrónicos agrupados
	req, err := http.NewRequest("POST", url, strings.NewReader(strings.Join(batch, "\n")))
	if err != nil {
		log.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")
	getenv.SetBasicAuth(req)

	client := &http.Client{}

	// Realizar la solicitud utilizando el cliente HTTP reutilizado
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error en la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

}
