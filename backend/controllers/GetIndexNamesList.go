package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func GetIndexNamesList(w http.ResponseWriter, r *http.Request) {
	url := getenv.GetZincSearchServerURL() + "api/index_name"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error creando la solicitud HTTP: %v", err)
	}

	getenv.SetBasicAuth(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error realizando la solicitud HTTP: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("la solicitud HTTP falló con código de estado %d", res.StatusCode)
	}

	var indexNames []string
	if err := json.NewDecoder(res.Body).Decode(&indexNames); err != nil {
		fmt.Printf("error decodificando la respuesta JSON: %v", err)
	}

	json.NewEncoder(w).Encode(indexNames)

}
