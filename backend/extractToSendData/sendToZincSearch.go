package data

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var client = &http.Client{} // Definir el cliente HTTP fuera de la función

// func SendToZincSearch(data []string, batchSize int) {
// 248.3292ms
// 	zincUser := "admin"
// 	zincPass := "Complexpass#123"
// 	url := "http://localhost:4080/api/_bulk"

// 	//divide el slice data en lotes de tamaño batchSize para procesar cada lote por separado
// 	for i := 0; i < len(data); i += batchSize {
// 		end := i + batchSize
// 		if end > len(data) {
// 			end = len(data)
// 		}

// 		// Agrupar los correos electrónicos en lotes
// 		batch := data[i:end]

// 		// Crear la solicitud HTTP con los correos electrónicos agrupados
// 		req, err := http.NewRequest("POST", url, strings.NewReader(strings.Join(batch, "\n")))
// 		if err != nil {
// 			fmt.Println("Error al crear la solicitud HTTP:", err)
// 			return
// 		}

// 		// Establecer las credenciales
// 		req.SetBasicAuth(zincUser, zincPass)

// 		// Establecer las cabeceras
// 		req.Header.Set("Content-Type", "application/json")

// 		// Realizar la solicitud utilizando el cliente HTTP reutilizado
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println("Error en la solicitud HTTP:", err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		// Imprimir la respuesta del servidor
// 		fmt.Println("Respuesta del servidor:", resp.Status)
// 	}
// }

// SendToZincSearch envía lotes de correos electrónicos a la API en paralelo.
func SendToZincSearch(data []string, batchSize int) {
	// 135.0969ms
	zincUser := "admin"
	zincPass := "Complexpass#123"
	url := "http://localhost:4080/api/_bulk"

	var wg sync.WaitGroup

	//divide el slice data en lotes de tamaño batchSize para procesar cada lote por separado
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		// Agrupar los correos electrónicos en lotes
		batch := data[i:end]

		// Aumentar el contador de WaitGroup para cada lote
		wg.Add(1)

		// Enviar el lote en una goroutine separada
		go SendBatch(batch, zincUser, zincPass, url, &wg)
	}

	// Esperar a que todas las goroutines finalicen
	wg.Wait()
}

// SendBatch envía un lote de correos electrónicos a la API en paralelo.
func SendBatch(batch []string, zincUser, zincPass, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Crear la solicitud HTTP con los correos electrónicos agrupados
	req, err := http.NewRequest("POST", url, strings.NewReader(strings.Join(batch, "\n")))
	if err != nil {
		fmt.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	// Establecer las credenciales
	req.SetBasicAuth(zincUser, zincPass)

	// Establecer las cabeceras
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud utilizando el cliente HTTP reutilizado
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error en la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Imprimir la respuesta del servidor
	fmt.Println("Respuesta del servidor:", resp.Status)
}
