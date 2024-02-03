package data

import (
	"fmt"
	"net/http"

	zinc "github.com/KerenBermeo/CorreoQuery/indexZincsearch"
)

// Función que envía datos a ZincSearch a través de una solicitud HTTP POST
func SendToZincSearch(data string) {
	// url para operacion masiva
	url := "http://localhost:4080/api/_bulk"

	client := &http.Client{}

	req, err := zinc.MakeRequestWithAuth("POST", url, data)
	if err != nil {
		fmt.Println("Error al pedir autorizacion:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error en la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respuesta del servidor:", resp.Status)
}
