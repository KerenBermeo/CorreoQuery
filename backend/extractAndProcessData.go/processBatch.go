package data

import "fmt"

func ProcessBatch(batch []byte, filePath string) ([]byte, error) {
	// Convierte el lote a una cadena para facilitar el análisis
	content := string(batch)

	jsonData, err := convertToJSON(content, filePath)
	if err != nil {
		return []byte{}, fmt.Errorf("error al pasar a json: %v", err)
	}

	return jsonData, nil
}
