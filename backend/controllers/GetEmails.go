package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
)

var NumberFiles int

func GetEmails(w http.ResponseWriter, r *http.Request) {

	serverURL := getenv.GetZincSearchServerURL()
	if serverURL == "" {
		log.Println("Variable de entorno vacia")
	}

	nameIndex := getenv.GetNameIndex()
	if nameIndex == "" {
		log.Println("Variable de entorno vacia")
	}

	url := fmt.Sprintf("%sapi/%s/_search/", serverURL, nameIndex)

	// Crear una nueva estructura de solicitud de búsqueda
	searchReq := model.SearchRequest{
		SearchType: "matchall",
		Query:      model.Query{},
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: NumberFiles,
		Source:     []string{},
	}

	// Codificar la estructura en JSON
	reqBody, err := json.Marshal(searchReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusInternalServerError)
		return
	}

	// Crear una nueva solicitud HTTP con el cuerpo codificado en JSON

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating HTTP request: %v", err), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Establecer la autenticación básica en la solicitud
	err = getenv.SetBasicAuth(req)
	if err != nil {
		log.Println("Error: ", err)
	}

	// Realizar la solicitud HTTP
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error performing HTTP request: %v", err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading response body: %v", err), http.StatusInternalServerError)
		return
	}

	// Escribir la respuesta en el cuerpo de la respuesta HTTP
	if _, err := w.Write(body); err != nil {
		http.Error(w, fmt.Sprintf("error writing response body: %v", err), http.StatusInternalServerError)
		return
	}
}
