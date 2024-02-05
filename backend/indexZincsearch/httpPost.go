package indexzincsearch

import (
	"fmt"
	"net/http"
	"strings"
)

// func httpPOST(url string, data []byte) {
// 	req, err := MakeRequestWithAuth("POST", url, data)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	log.Printf("Posting to: %s...", url)
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	log.Printf("Zinc server response code: %d", res.StatusCode)
// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	log.Printf("Zinc server response body: %s", string(resBody))
// }

func HttpPOST(url string, jsonBody string) {
	client := &http.Client{}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonBody))
	if err != nil {
		fmt.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	// Establecer las credenciales
	zincUser := "admin"
	zincPass := "Complexpass#123"
	req.SetBasicAuth(zincUser, zincPass)

	// Establecer las cabeceras
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error en la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Imprimir la respuesta del servidor
	fmt.Println("Respuesta del servidor:", resp.Status)
}
