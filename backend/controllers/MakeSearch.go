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

func MakeSearch(w http.ResponseWriter, r *http.Request) {
	// Obtener la URL del servidor de búsqueda desde la variable de entorno
	serverUrl := getenv.GetZincSearchServerURL()
	if serverUrl == "" {
		http.Error(w, "Variable de entorno vacía para la URL del servidor de búsqueda", http.StatusInternalServerError)
		return
	}

	nameIndex := getenv.GetNameIndex()
	if nameIndex == "" {
		log.Println("Variable de entorno vacia")
	}

	// Decodificar el cuerpo de la solicitud en una estructura
	var requestBody struct {
		Parameter string `json:"parameter"`
		Num       int    `json:"num"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decodificando el cuerpo de la solicitud: %v", err), http.StatusBadRequest)
		return
	}

	// Construir la URL para la solicitud de búsqueda
	url := fmt.Sprintf("%sapi/%s/_search", serverUrl, nameIndex)

	// Construir la solicitud de búsqueda
	searchReq := model.SearchRequest{
		SearchType: "matchphrase",
		Query: model.Query{
			Term:  requestBody.Parameter,
			Field: "_all",
		},
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: requestBody.Num,
		Source:     []string{},
	}

	// Codificar la estructura de la solicitud en JSON
	reqBody, err := json.Marshal(searchReq)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error codificando el cuerpo de la solicitud: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(reqBody))
	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creando la solicitud HTTP: %v", err), http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido de la solicitud
	req.Header.Set("Content-Type", "application/json")

	// Establecer la autenticación básica en la solicitud
	err = getenv.SetBasicAuth(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error estableciendo la autenticación básica: %v", err), http.StatusInternalServerError)
		return
	}

	// Realizar la solicitud HTTP
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error realizando la solicitud HTTP: %v", err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error leyendo el cuerpo de la respuesta: %v", err), http.StatusInternalServerError)
		return
	}

	// Escribir la respuesta en el cuerpo de la respuesta HTTP
	if _, err := w.Write(body); err != nil {
		http.Error(w, fmt.Sprintf("Error escribiendo el cuerpo de la respuesta: %v", err), http.StatusInternalServerError)
		return
	}
}
