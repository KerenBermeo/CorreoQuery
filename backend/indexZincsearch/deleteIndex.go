package indexzincsearch

import (
	"fmt"
	"net/http"
)

func DeleteIndex() {
	url := "http://localhost:4080/api/index/*"

	client := &http.Client{}
	req, err := MakeRequestWithAuth("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error al construir la solicitud de eliminación:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud de eliminación:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respuesta del servidor:", resp.Status)
}
