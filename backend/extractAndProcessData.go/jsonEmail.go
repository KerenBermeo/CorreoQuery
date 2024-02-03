package data

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/KerenBermeo/CorreoQuery/model"
)

// Función para procesar el contenido del archivo y convertirlo a JSON
func convertToJSON(content string, filePath string) ([]byte, error) {
	// Convierte el contenido del archivo a una cadena

	wordMap := map[string]struct{}{
		"Mime-Version":              {},
		"Content-Type":              {},
		"Content-Transfer-Encoding": {},
		"X-From":                    {},
		"X-To":                      {},
		"X-cc":                      {},
		"X-bcc":                     {},
		"X-Folder":                  {},
		"X-Origin":                  {},
		"X-FileName":                {},
	}

	fileContent := string(content)

	// Divide el contenido del archivo en líneas
	lines := strings.Split(fileContent, "\n")

	// Inicializa una estructura de Email para almacenar los datos
	email := model.Email{}

	// Itera sobre las líneas del archivo
	for _, line := range lines {
		// Divide cada línea en clave y valor
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// excluye las lineas con las palabras encontradas en el mapa
			if _, found := wordMap[key]; !found {
				switch key {
				// se agrega el contenido que tiene la llave correspondiente a la estructura
				case "Message-ID":
					email.MessageId = value
				case "From":
					email.From = value
				case "To":
					if value != "" {
						email.To = strings.Split(value, ",")
						for i := range email.To {
							email.To[i] = strings.TrimSpace(email.To[i])
						}
					} else {
						email.To = []string{}
						fmt.Printf("Campo 'To' vacío en el archivo: %s\n", filePath)
					}
				case "Subject":
					email.Subject = value

				case "Date":
					email.Date = value

				default:
					email.Content = email.Content + value
				}
			}

		}
	}

	// Convierte la estructura Email a JSON
	jsonData, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	return jsonData, err
}
