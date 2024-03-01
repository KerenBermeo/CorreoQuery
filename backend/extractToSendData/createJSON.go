package data

import (
	"encoding/json"
	"log"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/KerenBermeo/CorreoQuery/model"
)

func CreateJSON(emails []model.Email) {
	nameIndex := getenv.GetNameIndex()
	if nameIndex == "" {
		log.Println("Variable de entorno vacia")
	}

	bulk := model.Payload{
		Index:   nameIndex,
		Records: emails,
	}

	jsonData, err := json.Marshal(bulk)
	if err != nil {
		log.Println("Error al convertir a JSON:", err)
		return
	}

	SendData(jsonData)
}
