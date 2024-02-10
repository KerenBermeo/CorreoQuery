package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
	"github.com/go-chi/chi"
)

func SearchAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("_______________1_____________")
	index := chi.URLParam(r, "index_name")
	numStr := chi.URLParam(r, "num")
	fmt.Println("_______________donde para?_____________")
	// Convertir la cadena a un entero
	num, err := strconv.Atoi(numStr)
	if err != nil {
		// Manejar el error si la conversión falla
		http.Error(w, fmt.Sprintf("error converting 'num' parameter to integer: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("_______________1_____________")
	if index == "" {
		http.Error(w, "Index name is required", http.StatusBadRequest)
		return
	}

	fmt.Println("_______________2_____________")

	url := fmt.Sprintf("%s/api/%s/_search/", getenv.GetZincSearchServerURL(), index)

	// Crear una nueva estructura de solicitud de búsqueda
	searchReq := model.SearchRequest{
		SearchType: "matchall",
		Query:      struct{}{},
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: num,
		Source:     []string{},
	}

	// Codificar la estructura en JSON
	reqBody, err := json.Marshal(searchReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("_______________3_____________")
	fmt.Println("_______________Json_____________")
	fmt.Println(string(reqBody))
	fmt.Println("_______________Fin Json_____________")

	// Crear una nueva solicitud HTTP con el cuerpo codificado en JSON
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating HTTP request: %v", err), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Establecer la autenticación básica en la solicitud
	getenv.SetBasicAuth(req)

	// Realizar la solicitud HTTP
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error performing HTTP request: %v", err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Verificar el código de estado de la respuesta
	if res.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("HTTP request failed with status code %d", res.StatusCode), res.StatusCode)
		return
	}

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
