package controllers

import (
	"encoding/base64"
	"errors"
	"net/http"
)

func makeAPIRequest(typeOfRequest, url string) (*http.Response, error) {

	username := "admin"
	password := "Complexpass#123"

	// Crear la solicitud HTTP
	req, err := http.NewRequest(typeOfRequest, url, nil)
	if err != nil {
		return nil, err
	}

	// Agregar el encabezado de autorización a la solicitud
	req.Header.Set("Authorization", getAuthorizationHeader(username, password))

	// Realizar la solicitud HTTP
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Verificar el estado de la respuesta
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("API request failed with status code: " + res.Status)
	}

	return res, nil
}

func getAuthorizationHeader(username, password string) string {
	// Combina el nombre de usuario y la contraseña en una cadena
	auth := username + ":" + password
	// Codifica la cadena en base64
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// Agrega el prefijo "Basic " al encabezado de autorización
	return "Basic " + authEncoded
}
