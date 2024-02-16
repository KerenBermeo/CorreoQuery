package data

import (
	"encoding/json"
	"fmt"
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

func ConcurrentParsedEmailJson(parsedModels []model.Email) {
	fmt.Println(parsedModels)
	var batch []string

	for _, data := range parsedModels {

		jsonStr, err := ParsedEmailJson(data)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			continue
		}
		batch = append(batch, jsonStr)
	}

	sendBatch(batch)
}

func sendBatch(batch []string) {
	//fmt.Println(batch)

	// url := getenv.GetZincSearchServerURL() + "api/_bulk"

	// // Crear la solicitud HTTP con los correos electrónicos agrupados
	// req, err := http.NewRequest("POST", url, strings.NewReader(strings.Join(batch, "\n")))
	// if err != nil {
	// 	fmt.Println("Error al crear la solicitud HTTP:", err)
	// 	return
	// }

	// // Establecer el encabezado Content-Type
	// req.Header.Set("Content-Type", "application/json")
	// getenv.SetBasicAuth(req)

	// client := &http.Client{}

	// // Realizar la solicitud utilizando el cliente HTTP reutilizado
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error en la solicitud HTTP:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// // Imprimir la respuesta del servidor
	// fmt.Println("Respuesta del servidor:", resp.Status)

	// Incrementar el contador de lotes

}
