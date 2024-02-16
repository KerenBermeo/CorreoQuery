package controllers

import (
	"fmt"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func CheckIndexExists(nameIndex string) (bool, error) {
	url := getenv.GetZincSearchServerURL() + "api/index/" + nameIndex

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error creando la solicitud HTTP: %v", err)
		return false, err
	}

	getenv.SetBasicAuth(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error al enviar la solicitud HTTP: %v", err)
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
