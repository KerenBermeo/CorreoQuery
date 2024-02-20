package data

import (
	"log"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ConcurrentParsedEmailJson(parsedModels []model.Email) {

	var batch []string

	for _, data := range parsedModels {

		jsonStr, err := ParsedEmailJson(data)
		if err != nil {
			log.Println("Error al convertir a JSON:", err)
			continue
		}
		batch = append(batch, jsonStr)
	}

	SendBatch(batch)
}
