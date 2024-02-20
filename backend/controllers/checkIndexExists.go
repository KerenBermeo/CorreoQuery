package controllers

import (
	"fmt"
	"log"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func CheckIndexExists() (bool, error) {
	serverUrl := getenv.GetZincSearchServerURL()
	nameIndex := getenv.GetNameIndex()
	if nameIndex == "" {
		log.Println("Variable de entorno vacia")
	}

	if serverUrl == "" {
		log.Println("Variable de entorno vacia")
	}

	url := serverUrl + "api/index/" + nameIndex

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error creando la solicitud HTTP: %v", err)
		return false, err
	}

	err = getenv.SetBasicAuth(req)
	if err != nil {
		log.Println("Error: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error al enviar la solicitud HTTP: %v", err)
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return false, nil
	case http.StatusNotFound:
		return true, nil
	default:
		return false, fmt.Errorf("código de estado inesperado: %v", resp.StatusCode)
	}
}
