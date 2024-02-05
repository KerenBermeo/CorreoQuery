package data

import (
	"encoding/json"
	"log"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ParsedEmail(mail model.Email) (string, error) {
	emailJSON, err := json.Marshal(mail)
	if err != nil {
		log.Print("Error al convertir la estructura a JSON:", err)
		return "", err
	}

	return string(emailJSON), nil
}
