package data

import (
	"encoding/json"
	"log"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
)

// ParsedEmailJson convierte una estructura de correo electrónico (model.Email)
// en formato JSON y lo prepara para su indexación en Zincsearch.
func ParsedEmailJson(mail model.Email) (string, error) {
	nameIndex := getenv.GetNameIndex()
	if nameIndex == "" {
		log.Println("Variable de entorno vacia")
	}
	//indice a agregar
	strIndexName := `{ "index" : { "_index" : "` + nameIndex + `" } }` + "\n"

	emailJSON, err := json.Marshal(mail)
	if err != nil {
		log.Print("Error al convertir la estructura a JSON:", err)
		return "", err
	}
	// concatenar el indice con el json
	str := strIndexName + string(emailJSON)

	return str, nil
}
