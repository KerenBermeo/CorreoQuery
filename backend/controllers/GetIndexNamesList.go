package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func GetIndexNamesList(w http.ResponseWriter, r *http.Request) {
	serverURL := getenv.GetZincSearchServerURL()
	if serverURL == "" {
		log.Println("Variable de entorno vacia")
	}

	url := serverURL + "api/index_name"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error creando la solicitud HTTP: %v", err)
	}

	err = getenv.SetBasicAuth(req)
	if err != nil {
		log.Println("Error: ", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error realizando la solicitud HTTP: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("la solicitud HTTP falló con código de estado %d", res.StatusCode)
	}

	var indexNames []string
	if err := json.NewDecoder(res.Body).Decode(&indexNames); err != nil {
		log.Printf("error decodificando la respuesta JSON: %v", err)
	}

	json.NewEncoder(w).Encode(indexNames)

}
