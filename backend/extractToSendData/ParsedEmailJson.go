package data

import (
	"encoding/json"
	"log"

	"github.com/KerenBermeo/CorreoQuery/model"
)

// ParsedEmailJson convierte una estructura de correo electrónico (model.Email)
// en formato JSON y lo prepara para su indexación en Zincsearch.
func ParsedEmailJson(mail model.Email) (string, error) {
	//indice a agregar
	strIndexName := `{ "index" : { "_index" : "email" } }` + "\n"

	emailJSON, err := json.Marshal(mail)
	if err != nil {
		log.Print("Error al convertir la estructura a JSON:", err)
		return "", err
	}
	// concatenar el indice con el json
	str := strIndexName + string(emailJSON)

	return str, nil
}
