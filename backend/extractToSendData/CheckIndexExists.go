package data

import (
	"fmt"
	"net/http"
)

// CheckIndexExists verifica si un índice existe en la API.
func CheckIndexExists(indexName string) error {

	// Construir la URL para el índice específico
	url := fmt.Sprintf("http://localhost:4080/api/index/%s", indexName)

	// Crear una solicitud HTTP de tipo HEAD
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return err
	}

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Printf("El índice '%s' existe.\n", indexName)
		//Eliminar indice
		err := deleteIndex(indexName)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return nil
	case http.StatusNotFound:
		fmt.Printf("El índice '%s' no existe.\n", indexName)
		return nil
	default:
		return fmt.Errorf("error al verificar el índice. Código de estado: %d", resp.StatusCode)
	}
}

func deleteIndex(indexName string) error {
	// Construir la URL para el índice específico
	url := fmt.Sprintf("http://localhost:4080/api/index/%s", indexName)

	// Crear una solicitud HTTP de tipo DELETE
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("El índice '%s' ha sido eliminado.\n", indexName)
		return nil
	}

	return fmt.Errorf("error al eliminar el índice '%s'. Código de estado: %d", indexName, resp.StatusCode)
}
