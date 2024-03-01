package data

import (
	"bytes"
	"log"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func SendData(jsondata []byte) {
	serverUrl := getenv.GetZincSearchServerURL()
	if serverUrl == "" {
		log.Println("Variable de entorno vacia")
	}

	url := serverUrl + "api/_bulkv2"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsondata))
	if err != nil {
		log.Println("Error al crear la solicitud HTTP:", err)

	}
	req.Header.Set("Content-Type", "application/json")

	getenv.SetBasicAuth(req)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error en la solicitud HTTP:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Los archivos fueron indexados")
	}
}
